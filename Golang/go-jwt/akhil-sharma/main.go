package main

import (
	"jwt-golang-auth-akhil-sharma/database"
	"jwt-golang-auth-akhil-sharma/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load env variables

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading enviroment variables:", err)
	}

	// DB initialization
	database.DBinit()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	//routes
	routes.AuthRoutes(r)
	routes.UserRoutes(r)
	routes.PublicRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

}

func interpret(command string) string {

	var parser string
	var open bool
	for _, v := range command {
		switch {
		case v == '(':
			open = true

		case v == ')' && open:
			parser += "o"
		case v == ')' && !open:
			continue
		default:
			open = false
			parser += string(v)
		}

	}
	return parser
}
