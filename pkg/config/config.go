package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.AddConfigPath("./pkg/config")

	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
