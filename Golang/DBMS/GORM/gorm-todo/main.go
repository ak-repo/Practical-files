package main

import "github.com/gin-gonic/gin"

func init() {

	LoadEnvVariables()
	InitDB()
}

func main() {

	DB.AutoMigrate(&Todo{})

	router := gin.Default()

	TodoRoutes(router)

	router.Run()
}
