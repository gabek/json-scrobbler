package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config struct
type Config struct {
	LastFM          LastFM `yaml:"lastfm"`
	Input           string `yaml:"inputFile"`
	PollingInterval string `yaml:"pollingInterval"`
}

type LastFM struct {
	Key      string `yaml:"key"`
	Secret   string `yaml:"secret"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func getConfig() Config {
	filePath := "config/config.yaml"

	if !FileExists(filePath) {
		log.Fatal("ERROR: valid config/config.yaml is required")
	}

	yamlFile, err := ioutil.ReadFile(filePath)

	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	return config
}
