package v1

import (
	v1 "github.com/JLavrin/project-aigo/controller/path/v1"
	"github.com/gin-gonic/gin"
)

func Bind(r *gin.RouterGroup) {
	rg := r.Group("/v1/paths")

	// query
	rg.GET("", v1.Get)
}
