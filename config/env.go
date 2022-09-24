package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var EnvConf EnvConfig

type EnvConfig struct {
	BackendPort  string `yaml:"backend_port"`
	FrontendPort string `yaml:"frontend_port"`
}

func ParseEnvConfig() {
	file, err := ioutil.ReadFile(`config/env.yaml`)
	if err != nil {
		log.Fatal(`Events config file open error : `, err)
	}

	EnvConf = EnvConfig{}
	err = yaml.Unmarshal(file, &EnvConf)
	if err != nil {
		log.Fatal(`Cannot decode env configs : `, err)
	}
}
