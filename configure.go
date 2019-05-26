package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

//Configure is the data model for the configure YAML
type Configure struct {
	EC2 map[string]string `yaml:"EC2"`
	EBS map[string]string `yaml:"EBS"`
}

//GetConf iterates through the config Yaml
func (c *Configure) GetConf() *Configure {

	confFile, err := ioutil.ReadFile("gatfish.yaml")
	if err != nil {
		log.Printf("confFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(confFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
