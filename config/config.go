package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var C Config

func SetupConfig() {
	// Load YAML configuration
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Unmarshal YAML into the Config struct
	err = yaml.Unmarshal(data, &C)
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Config struct {
	Git struct {
		Command struct {
			Path string `yaml:"path"`
		} `yaml:"command"`
		Config struct {
			ConfigFilePath string `yaml:"configFilePath"`
		} `yaml:"gcloud"`
		Prompt struct {
			System   string `yaml:"system"`
			Template string `yaml:"template"`
		} `yaml:"prompt"`
		Validation struct {
			Pattern string `yaml:"pattern"`
		} `yaml:"validation"`
	}
}
