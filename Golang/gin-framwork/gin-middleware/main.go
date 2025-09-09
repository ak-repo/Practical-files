package main

import (
	"gin-middleware/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.New()
	router.Use(gin.Recovery(), middleware.LoggerMiddleware()) //apply for all routes

	//routes
	router.GET("/home", func(c *gin.Context) {
		c.String(http.StatusOK, "Homepage")
	})

	// auth middle ware only apply for admine,,, also can be used with Group routes.
	router.GET("/admin", middleware.Authenticate(), func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"messgae": "you are on admin page",
		})
	})

	//server
	router.Run()
}
