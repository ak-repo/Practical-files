package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(role string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		userRole := ctx.GetHeader("role")
		if userRole == "" {
			ctx.AbortWithStatusPureJSON(http.StatusUnauthorized, gin.H{"error": "no role specified"})
			return
		}

		if role == userRole {
			ctx.Next()
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "permission denied", "role": userRole})

	}
}
