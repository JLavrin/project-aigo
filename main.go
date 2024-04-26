package main

import (
	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		response := BaseResponse{
			Code:    200,
			Message: "Hello, World!",
			Data:    nil,
		}

		c.JSON(200, response)
	})

	err := router.Run(":8080")

	if err != nil {
		panic(err)
	}
}
