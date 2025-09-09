package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//authentication middleware
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "Bearer secret123" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorised entry",
			})
			return
		}

		c.Next()
	}
}



//logger middlerware
func LoggerMiddleware()gin.HandlerFunc{
	return func (c *gin.Context){

		//before request
		fmt.Println("Request path:", c.Request.URL.Path)

		c.Next()

		//after request
		status:= c.Writer.Status()
		fmt.Println("Response Status",status)
	}

}














