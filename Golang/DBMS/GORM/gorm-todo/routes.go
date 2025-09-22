package main

import "github.com/gin-gonic/gin"

func TodoRoutes(r *gin.Engine) {

	userGroup := r.Group("/todos")
	{

		// routes
		userGroup.POST("", TodoCreate)   // create
		userGroup.GET("/", TodoIndex)    // read all
		userGroup.GET("/:id", TodosShow) // read ome
		userGroup.PUT("/:id", TodoUpdate)    //update
		userGroup.DELETE("/:id", TodoDelete) //delete

	}
}
