package controllers

import (
	articleService "WriteWise/internal/modules/article/services"
	"WriteWise/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService articleService.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleService: articleService.New(),
	}
}

func (controller *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/html/home", gin.H{
		"title":    "Home page",
		"featured": controller.articleService.GetFeaturedArticle(),
		"stories":  controller.articleService.GetStoriesArticle(),
	})
}
