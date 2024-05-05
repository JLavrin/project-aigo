package v1

import (
	"github.com/JLavrin/project-aigo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PathRequest struct {
	StartX float64 `form:"startX"`
	StartY float64 `form:"startY"`
	EndX   float64 `form:"endX"`
	EndY   float64 `form:"endY"`
}

func Get(c *gin.Context) {
	var req PathRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := util.Request("/0.6/map.json", []string{"startX", "startY", "endX", "endY"})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var route any = nil

	c.JSON(http.StatusOK, gin.H{"route": route})
}
