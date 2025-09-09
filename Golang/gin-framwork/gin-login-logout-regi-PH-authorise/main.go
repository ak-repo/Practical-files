package main

import "github.com/gin-gonic/gin"

//type

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

// user database
var usersDB = map[string]Login{}

func main() {

	router := gin.Default()

	//routes
	//public
	public := router.Group("/")
	{
		public.GET("/products",ProductHandler)
	}

	//private
	private := router.Group("/user")
	{
		private.POST("/register", RegisterHandler)
		private.POST("/login", LoginHandler)
		private.POST("/logout", LogoutHandler)
		private.GET("/info", HandleProtectedRoute)
	}

	router.Run()
}
