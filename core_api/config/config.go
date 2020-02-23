package config

import (
	"github.com/spf13/viper"
	"os"
)

var Config struct {
	LogQueries bool
	LogLevel string
	LogCaller bool
	LogRequests bool
	AutoMigrate bool
	ApiBindAddress string
	DatabasePath string
}

func LoadConfig() error {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(configFile)
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		return err
	}
	return nil
}
