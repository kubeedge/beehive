package core

import (
	"time"

	"k8s.io/klog"

	"github.com/kubeedge/beehive/pkg/common/config"
	"github.com/kubeedge/beehive/pkg/core/context"
)

const (
	tryReadKeyTimes = 5
)

// Module interface
type Module interface {
	Name() string
	Group() string
	Start(c *context.Context)
	Cleanup()
}

var (
	// Modules map
	modules         map[string]Module
	disabledModules map[string]Module
)

func init() {
	modules = make(map[string]Module)
	disabledModules = make(map[string]Module)
	config.AddConfigChangeCallback(moduleChangeCallback{})
	eventListener := config.EventListener{Name: "eventListener1"}
	config.CONFIG.RegisterListener(eventListener, "modules.enabled")
}

// Register register module
func Register(m Module) {
	if isModuleEnabled(m.Name()) {
		modules[m.Name()] = m
		klog.Infof("module %s registered", m.Name())
	} else {
		disabledModules[m.Name()] = m
		klog.Infof("module %s is not register, please check modules.yaml", m.Name())
	}
}

//
func isModuleEnabled(m string) bool {
	modules := config.CONFIG.GetConfigurationByKey("modules.enabled")
	if modules != nil {

		for _, value := range modules.([]interface{}) {
			if m == value.(string) {
				return true
			}
		}
	}

	return false
}

type moduleChangeCallback struct{}

func (cb moduleChangeCallback) Callback(k string, v interface{}) {
	retryReadKey := func() interface{} {
		for times := 0; times < tryReadKeyTimes; times++ {
			// try to read the key again
			modules := config.CONFIG.GetConfigurationByKey(k)
			if modules != nil {
				return modules
			}
			time.Sleep(200 * time.Millisecond)
		}
		return nil
	}

	if k == "modules.enabled" {
		currentModules, ok := v.([]interface{})
		if !ok {
			klog.Infof("retry read key: %+v", k)
			modules := retryReadKey()
			if currentModules, ok = modules.([]interface{}); !ok {
				klog.Warningf("bad value of key(%s)", k)
				return
			}
		}

		newModules, deletedModules := calculateModuleChanges(currentModules)
		klog.Infof("current module list: %+v, disabledmodule: %+v addmodule: %+v  deletedmodule: %+v",
			currentModules, disabledModules, newModules, deletedModules)
		//Remove disabled modules
		for _, m := range deletedModules {
			module, exist := modules[m]
			if !exist {
				klog.Infof("module %s is not existing", m)
				break
			}
			module.Cleanup()
			delete(modules, m)
			disabledModules[m] = module
			klog.Infof("module %s is disabled", m)
		}
		//Enable new modules
		for _, m := range newModules {
			module := disabledModules[m]
			if module == nil {
				klog.Infof("module %s is not existing", m)
				break
			}
			Register(module)
			coreContext := context.GetContext(context.MsgCtxTypeChannel)
			//Init the module
			coreContext.AddModule(module.Name())
			//Assemble typeChannels for send2Group
			coreContext.AddModuleGroup(module.Name(), module.Group())
			go module.Start(coreContext)
			delete(disabledModules, m)
			klog.Infof("module %s is enabled", m)
		}
	}
}

func calculateModuleChanges(newModulesConfig []interface{}) ([]string, []string) {
	var newModules, deletedModules []string
	for module := range modules {
		if !inSlice(module, newModulesConfig) {
			deletedModules = append(deletedModules, module)
		}
	}
	for _, m := range newModulesConfig {
		if modules[m.(string)] == nil {
			newModules = append(newModules, m.(string))
		}
	}
	return newModules, deletedModules
}

func inSlice(e string, slice []interface{}) bool {
	for _, s := range slice {
		if s.(string) == e {
			return true
		}
	}
	return false
}

// GetModules gets modules map
func GetModules() map[string]Module {
	return modules
}
