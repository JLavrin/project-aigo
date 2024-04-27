package main

import (
	"github.com/JLavrin/project-aigo/config"
	"github.com/JLavrin/project-aigo/models"
	"github.com/JLavrin/project-aigo/routers"
	"net/http"
)

func main() {

	routerInit := routers.InitRouter()
	models.Setup()
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           config.Server.Port,
		Handler:        routerInit,
		ReadTimeout:    config.Server.ReadTimeout,
		WriteTimeout:   config.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
