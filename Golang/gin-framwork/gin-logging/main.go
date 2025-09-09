package main

import (
	middleware "gin-logging/middlerware"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// logging
	file, err := os.Create("gin-logging.log")
	if err != nil {
		panic(gin.ErrorLogger())
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	//define formate for log of routes in gin
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandler int) {
		log.Printf("Defines formatted log info: %v %v %v %v \n", httpMethod, absolutePath, handlerName, nuHandler)

	}

	//define formate of logs in GIn and file
	// Custom logger with formatter
	router.Use(gin.LoggerWithFormatter(middleware.CustomLogger))

	//routes
	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "Homepage",
		})
	})

	router.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Product list page",
		})
	})

	//server
	router.Run()
}
