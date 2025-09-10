package main

import "github.com/gin-gonic/gin"

func main() {

	//multiplexer
	router := gin.Default()

	//public routes
	router.GET("/", HomeHandler)
	router.GET("/posts", PostPageHandler)
	router.POST("/register", RegistrationHandler)
	router.POST("/login", LoginHandler)
	router.POST("/logout", LogoutHandler)

	// private routes

	private := router.Group("/user")
	private.Use(RequireAuth())
	{

		// post creation, updation, delete routes
		private.POST("/create", CreatePostHandler)
		private.PATCH("/update/:postid", RequireOwner(), UpdatePostHandler)
		private.DELETE("/delete/:postid", RequireOwner(), DeletePostHandler)

	}

	//server
	router.Run()

}
