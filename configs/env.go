package configs

import (
	"log"

	"github.com/engine/models"
	"github.com/spf13/viper"
)

func Environment() (c models.EnvConfig, err error) {
	viper.SetConfigName("env") // name of config file (without extension)
	viper.AddConfigPath(".")   // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found; ignore error if desired")
		} else {
			log.Println("Config file was found but another error was produced")
		}
	}

	viper.AutomaticEnv()

	// Config file found and successfully parsed
	err = viper.Unmarshal(&c)
	if err != nil {
		return
	}
	return
}
