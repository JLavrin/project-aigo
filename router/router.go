package router

import (
	articleV1 "github.com/JLavrin/project-aigo/router/article/v1"
	"github.com/JLavrin/project-aigo/src/domain/home/services"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, services.Get())
	})

	api := r.Group("/api")

	articleV1.Bind(api)

	return r
}
