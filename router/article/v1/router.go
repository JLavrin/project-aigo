package v1

import (
	v1 "github.com/JLavrin/project-aigo/controller/article/v1"
	"github.com/gin-gonic/gin"
)

func Bind(r *gin.RouterGroup) {
	rg := r.Group("/v1/articles")

	// query
	rg.GET("", v1.GetAll)
	rg.GET("/:id", v1.Get)

	// commands
	rg.POST("", v1.Create)
	rg.PUT("/:id", v1.Update)
	rg.DELETE("/:id", v1.Delete)
}
