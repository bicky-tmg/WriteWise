package controllers

import (
	"WriteWise/internal/modules/user/requests/auth"
	UserService "WriteWise/internal/modules/user/services"
	"log"

	"WriteWise/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
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
	user, err := ctrl.userService.Create(request)

	// Check if there is any error on the user creation
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	// after creating the user > redirect to homepage
	log.Printf("The user created successfully with a name %s \n", user.Name)
	c.Redirect(http.StatusFound, "/")
}
