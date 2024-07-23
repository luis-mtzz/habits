// pkg/config/config.go

// Package config provides functionality to load and parse the application configuration.
// It defines the structure of the configuration and offers a method to load it from a YAML file.

package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Discogs struct {
		APIKey string `yaml:"api_key"`
	} `yaml:"discogs"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
