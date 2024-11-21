package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Domain   string `yaml:"domain"`
	Interval int    `yaml:"interval"`
	AliDNS   struct {
		ApiKey    string `yaml:"api_key"`
		ApiSecret string `yaml:"api_secret"`
	} `yaml:"ali_dns"`
}

var GConfig Config

func Init(configFile string) {
	content, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(content, &GConfig)
	if err != nil {
		panic(err)
	}
}
