package config

import "go.uber.org/config"

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

func NewRedisConfig(provider config.Provider) (*RedisConfig, error) {
	var c RedisConfig
	if err := provider.Get("redis").Populate(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
