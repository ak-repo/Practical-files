package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// basic auth
	auth := gin.BasicAuth(gin.Accounts{
		"user1": "1234",
		"user2": os.Getenv("USER2_PASS"), // set in env variable
	})

	// admin routes
	admin := router.Group("/admin", auth)
	{
		admin.GET("/info", func(c *gin.Context) {
			user := c.MustGet(gin.AuthUserKey).(string)
			c.JSON(http.StatusOK, gin.H{
				"adminInfo": "You are logged in as " + user,
			})
		})
	}

	// public routes
	public := router.Group("/")
	{
		public.GET("/home", func(c *gin.Context) {
			c.String(http.StatusOK, "You are on the home page")
		})
	}

	// start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
