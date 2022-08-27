package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Influx Influx `yaml:"influx"`
	HmIP   HmIP   `yaml:"hmip"`
}

type Influx struct {
	Url          string `yaml:"url"`
	Token        string `yaml:"token"`
	Organization string `yaml:"organization"`
	Bucket       string `yaml:"bucket"`
}

type HmIP struct {
	AuthToken   string `yaml:"authToken"`
	AccessPoint string `yaml:"accessPoint"`
	UserAgent   string `yaml:"userAgent"`
}

func ReadConfig(file string) (*Config, error) {

	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}