package config

import (
	"go.uber.org/config"
)

type StanConfig struct {
	ClusterId string `yaml:"cluster_id"`
	ClientId  string `yaml:"client_id"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
}

func NewStanConfig(provider config.Provider) (*StanConfig, error) {
	var c StanConfig
	if err := provider.Get("nats-streaming").Populate(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
