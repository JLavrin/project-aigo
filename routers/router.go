package routers

import (
	"github.com/JLavrin/project-aigo/api/v1"
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

	apiV1 := r.Group("/api/v1")

	{
		apiV1.GET("/articles", v1.GetAllArticles)
		apiV1.GET("/articles/:id", v1.GetArticle)
		apiV1.POST("/article", v1.CreateArticle)
		apiV1.PUT("/article/:id", v1.UpdateArticle)
		apiV1.DELETE("/article/:id", v1.DeleteArticle)
	}

	return r
}
