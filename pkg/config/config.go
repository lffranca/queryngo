package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func New(pathString *string) (*Config, error) {
	pathConfig := "./queryingo.yaml"
	if pathString != nil {
		pathConfig = *pathString
	}

	data, err := ioutil.ReadFile(pathConfig)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

type Config struct {
	Servers []Server `json:"servers" yaml:"servers"`
}
