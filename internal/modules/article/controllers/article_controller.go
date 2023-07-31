package controllers

import (
	ArticleService "WriteWise/internal/modules/article/services"
	"WriteWise/pkg/html"
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
		html.Render(c, http.StatusInternalServerError, "templates/errors/html/500", gin.H{"title": "Server error", "message": "error while converting the id"})
		return
	}

	// Find the article from the database
	article, err := controller.articleService.Find(id)

	// If the article is not found, show error page
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/errors/html/404", gin.H{"title": "Page not found", "message": err.Error()})
		return
	}

	// If article found, render article template
	html.Render(c, http.StatusOK, "modules/article/html/show", gin.H{"title": "Show article", "article": article})
}
