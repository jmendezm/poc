package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var config *Configuration

type Configuration struct {
	Database      Database `yaml:"database"`
	ApiListenPort int      `yaml:"api_listen_port"`
}

type Database struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	User         string `yaml:"user"`
	Pass         string `yaml:"pass"`
	DatabaseName string `yaml:"db_name"`
}

func Init(configFilePath string) *Configuration {
	f, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal("reading configuration file: ", err.Error())
	}
	if err = yaml.Unmarshal(f, &config); err != nil {
		log.Fatal("parsing configuration file: ", err.Error())
	}
	return config
}

func Get() *Configuration {
	if config == nil {
		log.Fatal("nil config provider")
	}
	return config
}
