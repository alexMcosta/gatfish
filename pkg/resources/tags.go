package resources

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

//Configure is the data model for the configure YAML
type Tags struct {
	ALL map[string]string `yaml:"ALL"`
	EC2 map[string]string `yaml:"EC2"`
	EBS map[string]string `yaml:"EBS"`
}

//GetConf reads through the config Yaml
func (t *Tags) Configure() *Tags {

	confFile, err := ioutil.ReadFile("gatfish.yaml")
	if err != nil {
		log.Printf("Error: There is no gatfish.yaml err= %v", err)
	}

	err = yaml.Unmarshal(confFile, t)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return t
}
