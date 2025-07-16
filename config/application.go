package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type ApplicationConfiguration struct {
	PostgresDsn     string `mapstructure:"POSTGRES_DSN"`
	ApplicationHost string `mapstructure:"APPLICATION_HOST"`
	ApplicationPort string `mapstructure:"APPLICATION_PORT"`
}

var AppConfig *ApplicationConfiguration

func InitializeCConfiguration() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.AutomaticEnv()
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}
}
