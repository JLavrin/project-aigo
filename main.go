package main

import (
	"github.com/JLavrin/project-aigo/config"
	"github.com/JLavrin/project-aigo/db"
	"github.com/JLavrin/project-aigo/router"
	"github.com/JLavrin/project-aigo/util"
	"net/http"
)

func main() {
	util.LoadEnvs()

	db.Setup()

	routerInit := router.InitRouter()
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           "127.0.0.1:" + config.Server.Port,
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
