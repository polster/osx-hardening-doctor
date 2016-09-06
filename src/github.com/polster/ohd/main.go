package main

import (
	"flag"
	"fmt"
	"time"
)

// App version
var Version = "0.3"

func main() {

	var commandFile string

	version := flag.Bool("version", false, "Prints the script's version and exits")
	flag.StringVar(&commandFile, "config_file", "checks.yml", "YAML file containing the commands and configuration")

	flag.Parse()

	// Print the script's version and exit
	if(*version) {
		fmt.Printf("OSX Hardening Doctor: %s\n", Version)
		return
	}

	DoOsVersionCheck()

	// Read the config file and quit on error
	err := ReadConfigRules(commandFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run commands and print results
	ruleCount := 0
	failCount := 0

	for _, rule := range ConfigRules {
		if rule.Enabled {
			checkCommand := rule.CheckCommand
			ruleCount++

			result := RunCommand(checkCommand)

			resultText := "\033[32mPASSED\033[39m"
			if !result {

				failCount++
				resultText = "\033[31mFAILED\033[39m"

			}

			fmt.Printf("[%s] %s\n", resultText, rule.Title)

		}
	}

	// Print summary
	fmt.Printf("-------------------------------------------------------------------------------\n")
	fmt.Printf("OSX Hardening Doctor: %s\n", Version)
	t := time.Now()
	fmt.Printf("Date: %s\n", t.Format("2006-01-02T15:04:05-07:00"))
	sysinfo := GetSystemInfo()
	fmt.Printf("SerialNumber: %s\nHardwareUUID: %s\n", sysinfo.SerialNumber, sysinfo.HardwareUUID)
	fmt.Printf("Final Score %d%%; Pass rate: %d/%d\n",
	CalculateScore(ruleCount, failCount),
	(ruleCount-failCount), ruleCount)
}