package config

import "github.com/spf13/viper"

var Configuration Config

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Configuration)
	if err != nil {
		panic(err)
	}
}
