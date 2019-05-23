package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

func YamlPath() string {
	if runtime.GOOS == "windows" {
		return "C:\\ProgramData\\PuppetLabs\\puppet\\cache\\state\\last_run_report.yaml"
	}

	return "/opt/puppetlabs/puppet/cache/state/last_run_report.yaml"
	// This is for when testing with a local report file
	// filename, _ := filepath.Abs("./last_run_report.yaml")
	// fmt.Print(filename)
	// return filename
}

func GetPuppetYaml() (*PuppetInfo, error) {

	var yamlData PuppetInfo
	yamlFile, err := os.Open(YamlPath())

	if err != nil {
		fmt.Print(err)
		return &yamlData, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(yamlFile)
	yamlString := buf.String()
	yamlString = strings.Replace(yamlString, "\r", "\n", -1)

	err = yaml.Unmarshal([]byte(yamlString), &yamlData)
	if err != nil {
		fmt.Print("Error during unmarshal")
		fmt.Print(err)
		return &yamlData, err
	}

	return &yamlData, nil
}
