package middleware

import (
	"log"
	jwtpkg "machine-task/pkg/jwt_pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMIddleWare(SecrectKey string) gin.HandlerFunc {

	return func(c *gin.Context) {

		// get token
		parts := c.GetHeader("Authorization")
		if parts == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
			return
		}

		token := strings.Split(parts, " ")[1]
		log.Println(token)

		// verify
		if err := jwtpkg.ValidateToken(SecrectKey, token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Next()

	}
}
