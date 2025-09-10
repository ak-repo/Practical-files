package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authentication middleware
func RequireAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		sToken, err := c.Cookie("session_token")
		if err != nil || sToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": " please login first"})
			c.Abort()
			return
		}
		// check the user id
		var currentUser *User
		for _, u := range UsersDB {
			if u.SessionToken == sToken {
				currentUser = &u
				break
			}
		}

		if currentUser == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}

		c.Set("user", currentUser.Username)
		c.Next()
	}

}

// Authorization for delete , update

func RequireOwner() gin.HandlerFunc {

	return func(c *gin.Context) {

		user := c.MustGet("user").(string)
		postID := c.Param("postid")

		post, ok := PostList[postID]
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
			c.Abort()
			return
		}

		if post.Auther != user {
			c.JSON(http.StatusForbidden, gin.H{"error": "not allowed to modify this post"})
			c.Abort()
			return
		}

		c.Next()

	}
}
