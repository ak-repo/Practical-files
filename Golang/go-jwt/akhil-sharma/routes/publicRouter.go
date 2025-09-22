package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(routes *gin.Engine){

	routes.GET("/",func (c *gin.Context){

		c.String(http.StatusOK,"Home page")
	})

}
