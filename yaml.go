package main

import (
	"fmt"
	"os"
	"runtime"

	"gopkg.in/yaml.v3"
	// "github.com/ghodss/yaml"
)

func YamlPath() string {
	if runtime.GOOS == "windows" {
		return "C:\\ProgramData\\PuppetLabs\\puppet\\cache\\state\\last_run_report.yaml"
	}

	return "/opt/puppetlabs/puppet/cache/state/last_run_report.yaml"
}

func GetPuppetYaml() (*PuppetInfo, error) {

	var yamlData PuppetInfo

	yamlFile, err := os.Open(YamlPath())
	if err != nil {
		fmt.Println(err)
		return &yamlData, err
	}
	defer yamlFile.Close()
	if err := yaml.NewDecoder(yamlFile).Decode(&yamlData); err != nil {
		return &yamlData, err
	}
	return &yamlData, nil
}
