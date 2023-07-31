package controllers

import (
	"WriteWise/internal/modules/user/requests/auth"
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
	// validate the request
	var request auth.RegisterRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}
	// Create the user

	// Check if there is any error on the user creation

	// after creating the user > redirect to homepage
	c.JSON(http.StatusOK, gin.H{"message": "Register done..."})
}
