package main

import (
	"gopkg.in/yaml.v2"
	"os/exec"
	"strings"
	"path/filepath"
	"fmt"
	"os"
	"io/ioutil"
)

// ConfigRules holds our yaml file containing our config
var ConfigRules ConfigRuleList

// ConfigRule is a container for each individual rule
type ConfigRule struct {
	Title        string            `yaml:"title"`
	CheckCommand string            `yaml:"check_command"`
	Enabled      bool              `yaml:"enabled"`
}

// ConfigRuleList is an array
type ConfigRuleList []ConfigRule

// ReadFile takes a relative path and returns the bytes in that file
func ReadFile(filename string) (data []byte, err error) {
	path, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println("ERROR: Unable to file:", filename)
		return nil, err
	}

	filehandle, err := os.Open(path)
	if err != nil {
		fmt.Println("ERROR: Error opening file:", path)
		return nil, err
	}
	defer filehandle.Close()

	data, err = ioutil.ReadAll(filehandle)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ReadConfigRules reads our yaml file
func ReadConfigRules(configFile string) error {
	ruleFile, err := ReadFile(configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(ruleFile, &ConfigRules)
	if err != nil {
		return err
	}
	return nil
}

// RunCommand returns true if the audit passed, or command was successful
func RunCommand(cmd string) bool {
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return false
	}

	return true
}

// GetCommandOutput runs a command and returns it's output
func GetCommandOutput(cmd string) string {
	out, _ := exec.Command("bash", "-c", cmd).Output()
	return strings.TrimSpace(string(out))
}

// DoOsVersionCheck performs OS version compatibility check
func DoOsVersionCheck() {

	osVersion := GetCommandOutput("system_profiler SPSoftwareDataType | grep \"System Version\" | cut -d: -f2")
	if !strings.Contains(osVersion, "OS X 10.11") {
		fmt.Println("ERROR: Unsupported OS. This tool was meant to be used only on OSX 10.11 (El Capitan)")
		return
	}
}
