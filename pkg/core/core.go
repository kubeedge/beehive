/*
Copyright 2019 The KubeEdge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package core is
package core

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kubeedge/beehive/pkg/common/log"
	"github.com/kubeedge/beehive/pkg/core/context"
)

// StartModules starts modules that are registered
func StartModules() {
	coreContext := context.GetContext(context.MsgCtxTypeChannel)

	modules := GetModules()
	for name, module := range modules {
		//Init the module
		coreContext.AddModule(name)
		//Assemble typeChannels for send2Group
		coreContext.AddModuleGroup(name, module.Group())
		go module.Start(coreContext)
		log.LOGGER.Info("starting module " + name)
	}
}

// GracefulShutdown is if it gets the special signals it does modules cleanup
func GracefulShutdown() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP, syscall.SIGABRT)
	select {
	case s := <-c:
		log.LOGGER.Info("got os signal " + s.String())
		//Cleanup each modules
		modules := GetModules()
		for name, module := range modules {
			log.LOGGER.Info("Cleanup module " + name)
			module.Cleanup()
		}
	}
}

//Run starts the modules and in the end does module cleanup
func Run() {
	//Address the module registration and start the core
	StartModules()
	// monitor system signal and shutdown gracefully
	GracefulShutdown()
}
