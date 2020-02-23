package main

import (
	" reqyt.run/worker/container_manager"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config = struct{
	Images container_manager.Config
}

func readConfig() Config {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal("no config file")
	}
	config := Config{}
	viper.SetConfigType("yaml")
	err = viper.ReadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	viper.Unmarshal(&config)
	return config
}
