package main

import "github.com/spf13/viper"

type Config interface {
}

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("configs")
	var err = viper.ReadInConfig()
	if err != nil {
		log.Criticalf ("Cannot read config file: %s",err)
		panic("Cannot read config file")
	}
}
