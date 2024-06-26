package util

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Envs struct {
	Port        string `mapstructure:"PORT" envDefault:"8080"`
	OpenAiToken string `mapstructure:"OPEN"`
}

func LoadEnvs() Envs {
	var result map[string]interface{}
	var envs Envs

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&result); err != nil {
		panic(err)
	}

	if err := mapstructure.Decode(&result, &envs); err != nil {
		panic(err)
	}

	return envs
}
