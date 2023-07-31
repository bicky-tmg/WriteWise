package controllers

import (
	"WriteWise/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/html/register", gin.H{"title": "Register"})
}

func (ctrl *Controller) HandleRegister(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register done..."})
}
