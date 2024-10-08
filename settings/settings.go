package settings

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed settings.yaml
var settingsFile []byte

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Settings struct {
	DB DatabaseConfig `yaml:"database"`
}

func New() (*Settings, error) {
	var s Settings

	if err := yaml.Unmarshal(settingsFile, &s); err != nil {
		return nil, err
	}

	return &s, nil
}
