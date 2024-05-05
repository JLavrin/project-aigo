package main

import (
	"github.com/JLavrin/project-aigo/config"
	"github.com/JLavrin/project-aigo/db"
	"github.com/JLavrin/project-aigo/router"
	"github.com/spf13/viper"
	"net/http"
)

type envConfigs struct {
	OpenAiToken string `mapstructure:"OPEN_AI_TOKEN"`
}

func main() {
	viper.SetConfigType("env")
	viper.AddConfigPath(".env")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	db.Setup()

	routerInit := router.InitRouter()
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           config.Server.Port,
		Handler:        routerInit,
		ReadTimeout:    config.Server.ReadTimeout,
		WriteTimeout:   config.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	err = server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
