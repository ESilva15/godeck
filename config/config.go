package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
)

type Authentication struct {
	User        string `yaml:"user"`
	Pass        string `yaml:"pass"`
	AppPassword string `yaml:"apppassword"`
}

type Configuration struct {
	Auth Authentication `yaml:"authentication"`
	Url  string         `yaml:"url"`
}

var instance *Configuration
var once sync.Once

func GetInstance() *Configuration {
	once.Do(func() {
		instance = &Configuration{}
		instance.loadConfiguration()
	})

	return instance
}

func (c *Configuration) loadConfiguration() {
	confPath := "./config.yaml"

	file, err := os.ReadFile(confPath)
	if err != nil {
		log.Fatalf("Unable to open configuration file [%s]: %s", confPath,
			err.Error())
	}

	err = yaml.Unmarshal(file, &instance)
	if err != nil {
		log.Fatalf("Error parsing config file: %s", err.Error())
	}
}
