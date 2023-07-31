package routes

import (
	articleRoutes "WriteWise/internal/modules/article/routes"
	homeRoutes "WriteWise/internal/modules/home/routes"
	userRoutes "WriteWise/internal/modules/user/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoutes.Routes(router)
	articleRoutes.Routes(router)
	userRoutes.Routes(router)
}
