package api

import (
	"fmt"
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	Common *CommonConfig `mapstructure:"common"`
	Logger *LoggerConfig `mapstructure:"logger"`
	Server *ServerConfig `mapstructure:"server"`
}

type CommonConfig struct {
	Debug bool `mapstructure:"debug"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

func NewConfig(filePath string) (*Config, error) {
	configName := path.Base(filePath)
	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	configDir := path.Dir(filePath)
	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	return &Config{
		Common: &CommonConfig{
			Debug: viper.GetBool("common.debug"),
		},
		Logger: &LoggerConfig{
			Level: viper.GetString("logger.level"),
		},
		Server: &ServerConfig{
			Port: viper.GetInt("server.port"),
		},
	}, nil
}
