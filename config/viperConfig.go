package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type AppConfig struct {
	TelegramAPI     TelegramConfig
	URLShortenerAPI URLShortenerConfig
	Bolt            BoltConfig
}

type TelegramConfig struct {
	Token string
}

type URLShortenerConfig struct {
	ServerHost string
	ServerPort string
	SSL        bool
}

type BoltConfig struct {
	DbPath string
	Mode   os.FileMode
}

type ConfigDriver struct {
	v *viper.Viper
}

func LoadNewConfig() (*ConfigDriver, error) {
	v := viper.New()
	v.SetConfigFile(findConfigPath())

	if err := v.ReadInConfig(); err != nil {
		return nil, errors.New("unable to find config, create at least `config-local.yml`")
	}
	return &ConfigDriver{
		v: v,
	}, nil
}

func (c *ConfigDriver) ParseConfig() (AppConfig, error) {
	config := AppConfig{}

	err := c.v.Unmarshal(&config)
	if err != nil {
		return AppConfig{}, fmt.Errorf("config: unable to decode into struct: %w", err)
	}
	return config, nil
}

func findConfigPath() string {
	configPaths := map[string]string{
		"docker": "config/config-docker.yml",
		"local":  "config/config-local.yml",
	}

	pathKey := os.Getenv("CONFIG")
	if pathKey == "" {
		return configPaths["local"]
	}
	return configPaths[pathKey]
}
