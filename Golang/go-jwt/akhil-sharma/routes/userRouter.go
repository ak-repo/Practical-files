package routes

import (
	"jwt-golang-auth-akhil-sharma/controllers"
	"jwt-golang-auth-akhil-sharma/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	router.Use(middleware.Authenticate())
	router.GET("/users",controllers.GetUsers)
	router.GET("/users/:user_id",controllers.GetUser)
	

}
