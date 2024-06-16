package config

import (
	"go.uber.org/config"
)

type PgConfig struct {
	Name        string `yaml:"name"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	SslMode     string `yaml:"ssl_mode"`
	DsnTemplate string `yaml:"dsn_template"`
}

func NewPgConfig(provider config.Provider) (*PgConfig, error) {
	var c PgConfig
	if err := provider.Get("pg").Populate(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
