package main

import (
	config "github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	APIAddress string
}

func CreateConfig() (*Config, error) {
	c := Config{}

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadExists("config.default.yml", "config.yml")
	if err != nil {
		return &c, err
	}

	config.LoadOSEnv([]string{}, false)

	err = config.BindStruct("", &c)

	return &c, err
}