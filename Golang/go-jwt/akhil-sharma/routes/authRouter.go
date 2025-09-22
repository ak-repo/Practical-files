package routes

import (
	"jwt-golang-auth-akhil-sharma/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {

	auth := router.Group("/user")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}
}
