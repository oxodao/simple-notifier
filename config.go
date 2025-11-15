package main

import (
	"errors"
	"os"

	"github.com/oxodao/simple-notifier/notification_service"
	"gopkg.in/yaml.v3"
)

const CONFIG_PATH = "/etc/simple_notifier.yaml"

type RawLocation struct {
	Type     string         `yaml:"type"`
	Settings map[string]any `yaml:"settings"`
}

type Config struct {
	RawLocations map[string]RawLocation                   `yaml:"locations"`
	Locations    map[string]notification_service.Location `yaml:"-"`
}

func LoadConfig() (*Config, error) {
	if _, err := os.Stat(CONFIG_PATH); os.IsNotExist(err) {
		return nil, errors.New("failed to find config file at " + CONFIG_PATH)
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

	cdg.Locations = make(map[string]notification_service.Location)
	for k, v := range cdg.RawLocations {
		// Re-marshal settings
		// Meh but oh well
		settingsData, err := yaml.Marshal(v.Settings)
		if err != nil {
			return nil, err
		}

		if parseFunc, ok := notification_service.KNOWN_LOCATIONS[v.Type]; ok {
			loc, err := parseFunc(settingsData)
			if err != nil {
				return nil, err
			}

			cdg.Locations[k] = loc
		} else {
			return nil, errors.New("Unknown location type " + v.Type)
		}
	}

	return &cdg, nil
}
