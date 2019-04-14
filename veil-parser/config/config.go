package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var Config config

type config struct {
	Debug bool `yaml:"debug"`
	Redis struct {
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
	} `yaml:"redis"`
}

func LoadFromFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &Config)
}
