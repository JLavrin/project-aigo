package routers

import (
	"github.com/JLavrin/project-aigo/service/hello"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, hello.Get())
	})

	//apiV1 := r.Group("/api/v1")

	{

	}
	return r
}
