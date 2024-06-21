package config

import (
	"go.uber.org/config"
	"go.uber.org/fx"
)

const (
	EnvLocal = "config/.local.yaml"
	EnvDev   = "config/.dev.yaml"
	EnvProd  = "config/.prod.yaml"
)

type BaseConfig struct {
	Name string `yaml:"name"`
}

type Config struct {
	fx.Out
	Provider config.Provider
	Config   BaseConfig
}

func New(stage string) (Config, error) {
	var configName string

	switch stage {
	case "local":
		configName = EnvLocal
	case "dev":
		configName = EnvDev
	default:
		configName = EnvLocal
	}

	loader, err := config.NewYAML(config.File(configName))
	if err != nil {
		return Config{}, err
	}

	cfg := BaseConfig{
		Name: "default",
	}
	if err = loader.Get("app").Populate(&cfg); err != nil {
		return Config{}, err
	}

	return Config{Provider: loader}, nil
}
