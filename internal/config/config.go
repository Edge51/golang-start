package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig()  {
	viper.SetConfigFile("config/config/yaml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}