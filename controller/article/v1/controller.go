package v1

import (
	"github.com/JLavrin/project-aigo/src/domain/article/model"
	"github.com/JLavrin/project-aigo/src/domain/article/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

func Create(c *gin.Context) {
	article := model.Article{}

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	if err := validate.Struct(article); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id, err := service.Create(&article)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}

func Get(c *gin.Context) {
	var id int
	id, _ = strconv.Atoi(c.Param("id"))

	article, err := service.Get(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, article)
}

func GetAll(c *gin.Context) {
	articles, err := service.GetAll()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, articles)
}

func Update(c *gin.Context) {
	var id int
	id, _ = strconv.Atoi(c.Param("id"))
	article := model.Article{}

	validate := validator.New()

	if err := validate.Struct(article); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := service.Update(&article, &id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": article.ID})
}

func Delete(c *gin.Context) {
	var id int
	id, _ = strconv.Atoi(c.Param("id"))

	err := service.Delete(&id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"id": id})
}
