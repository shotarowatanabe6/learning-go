package api

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Common *CommonConfig `mapstructure:"common"`
	Server *ServerConfig `mapstructure:"server"`
}

type CommonConfig struct {
	Debug bool `mapstructure:"debug"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

func NewConfig(filePath string) (*Config, error) {
	viper.SetConfigName(filePath)
	viper.SetConfigType("toml")
	viper.AddConfigPath(filePath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	return &Config{
		Common: &CommonConfig{
			Debug: viper.GetBool("common.debug"),
		},
		Server: &ServerConfig{
			Port: viper.GetInt("server.port"),
		},
	}, nil
}
