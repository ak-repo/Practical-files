package main

import (
	"log"
	"machine-task/config"
	"machine-task/internals/handler"
	"machine-task/internals/middleware"
	"machine-task/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {

	// env
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error while laodinf env ", err)
	}

	// DB
	err = db.NewDB(cfg.DSN)

	if err != nil {
		log.Fatal("error while db connection")
	}

	// routes
	r := gin.Default()

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)

	protected := r.Group("/user")
	protected.Use(middleware.AuthMIddleWare(cfg.SecrectKey))
	{
		protected.GET("/logout", handler.Logout)
	}

	r.Run()
}
