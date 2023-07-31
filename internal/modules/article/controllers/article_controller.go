package controllers

import (
	ArticleService "WriteWise/internal/modules/article/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService ArticleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	// Get the article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "error while converting the id"})
		return
	}

	// Find the article
	article, err := controller.articleService.Find(id)

	// If the article is not found, show error page
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	// If article found, render article template
	c.JSON(http.StatusOK, gin.H{"article": article})
}
