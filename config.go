package main

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIG_PATH = "/etc/simple_notifier.yaml"

type Locations struct {
	Type string `yaml:"type"`
	Webhook string `yaml:"webhook"`
	BotName string `yaml:"bot_name"`
}

type Config struct {
	Locations map[string]Locations `yaml:"locations"`
}

func LoadConfig() (*Config, error) {
	if _, err := os.Stat(CONFIG_PATH); os.IsNotExist(err) {
		return nil, errors.New("Failed to find config file at " + CONFIG_PATH)
	}

	data, err := os.ReadFile(CONFIG_PATH)
	if err != nil {
		return nil, err
	}

	var cdg Config

	err = yaml.Unmarshal(data, &cdg)
	if err != nil {
		return nil, err
	}

	return &cdg, nil
}
