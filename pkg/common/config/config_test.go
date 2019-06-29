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

// Package config_test is
package config_test

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	archaius "github.com/go-chassis/go-archaius"

	"github.com/kubeedge/beehive/pkg/common/config"
	"github.com/kubeedge/beehive/pkg/common/util"
)

func TestInitializeConfig(t *testing.T) {
	configDir := util.GetCurrentDirectory() + "/conf"
	//create if not exists
	_, err := os.Stat(configDir)
	if !os.IsExist(err) {
		os.Mkdir(configDir, os.ModePerm)
	}
	defer os.RemoveAll(configDir)
	err = prepareConfigFile(configDir)
	addSources(configDir)
	if err != nil {
		t.Error(err)
	}

	loggerLevel := config.GetString("loggerLevel", "ERROR")
	if loggerLevel != "DEBUG" {
		t.Error("config info incorrect")
	}
	isEnabled := isModuleEnabled("eventbus")
	if !isEnabled {
		t.Error("Error to get modules enabled information")
	}
}

func isModuleEnabled(m string) bool {
	modules := config.Get("modules.enabled")
	if modules == nil {
		return false
	}
	for _, value := range modules.([]interface{}) {
		if m == value.(string) {
			return true
		}
	}
	return false
}

func prepareConfigFile(dir string) error {
	//Write log config file
	logConfigFile := dir + "/logging.yaml"
	logConfigContent := "loggerLevel: \"DEBUG\"\n" +
		"enableRsyslog: false\n" +
		"logFormatText: true\n" +
		"writers: [stdout]"
	err := writeConfigFile(logConfigContent, logConfigFile)
	if err != nil {
		return err
	}
	//Write module config file
	moduleConfigFile := dir + "/modules.yaml"
	moduleConfigContent := "modules:\n" +
		"  enabled: [eventbus, servicebus]"
	err = writeConfigFile(moduleConfigContent, moduleConfigFile)
	if err != nil {
		return err
	}
	return nil
}

func addSources(location string) error {
	err := filepath.Walk(location, func(location string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		ext := strings.ToLower(path.Ext(location))
		if ext == ".yml" || ext == ".yaml" {
			archaius.AddFile(location)
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func writeConfigFile(content string, fileName string) error {
	//Prepare log config file
	//Delete it if already exists
	if isFileExists(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			return err
		}
	} else { //Create file and fill the content
		f, err := os.Create(fileName)
		defer f.Close()
		if err != nil {
			return err
		}
		_, err = io.WriteString(f, content)
		if err != nil {
			return err
		}
	}
	return nil
}

func isFileExists(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
