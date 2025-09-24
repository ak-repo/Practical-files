package middleware

import (
	"jwt-gin-gorm/pkg/jwt_pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Verify JWT token for incomming request

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		//Get authorization Header
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		// Check Bearer scheme
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization"})
			return
		}

		tokenString := parts[1]

		// Parse and validate token
		claims, err := jwt_pkg.ValidateToken(tokenString, jwtSecret)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// set user info in context

		ctx.Set("user_id", claims.UserID)
		ctx.Set("email", claims.Email)
		ctx.Set("role", claims.Role)

		ctx.Next()

	}
}
