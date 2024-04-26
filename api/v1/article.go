package v1

import (
	"fmt"
	"github.com/JLavrin/project-aigo/service/article_service"
	"github.com/gin-gonic/gin"
)

/**
 * @Summary Create article
 * @Produce json
 * @Success 200 {object} Article
 * @Router /api/v1/article [post]
 */
func CreateArticle(c *gin.Context) {
	as := article_service.Article{}

	c.JSON(200, as.Create())
}

/**
 * @Summary Get article
 * @Produce json
 * @Param id path int true "Article ID"
 * @Success 200 {object} Article
 */
func GetArticle(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	as := article_service.Article{}

	c.JSON(200, as.Get())
}

/**
 * @Summary Get all articles
 * @Produce json
 * @Success 200 {object} Article
 * @Router /api/v1/articles [get]
 */
func GetAllArticles(c *gin.Context) {
	as := article_service.Article{}

	c.JSON(200, as.GetAll())
}

/**
 * @Summary Update article
 * @Produce json
 * @Success 200 {object} Article
 * @Router /api/v1/article [put]
 */
func UpdateArticle(c *gin.Context) {
	as := article_service.Article{}

	c.JSON(200, as.Update())
}

/**
 * @Summary Delete article
 * @Produce json
 * @Success 200 {object} Article
 * @Router /api/v1/article [delete]
 */
func DeleteArticle(c *gin.Context) {
	as := article_service.Article{}

	c.JSON(200, as.Delete())
}
