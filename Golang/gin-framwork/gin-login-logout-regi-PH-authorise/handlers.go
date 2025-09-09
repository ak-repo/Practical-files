package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// register
func RegisterHandler(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	// check vaild formate
	if len(username) < 4 || len(password) < 8 {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Not acceptable details"})
		return
	}

	//checking username is already used
	if _, ok := usersDB[username]; ok {
		c.JSON(http.StatusConflict, gin.H{"error": "username already taken"})
		return
	}

	//saving
	hashedPassword, _ := HashPassword(password)
	usersDB[username] = Login{HashedPassword: hashedPassword}

	c.JSON(http.StatusOK, gin.H{"success": "Registration completed"})

}

//login

func LoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	// varify
	user, ok := usersDB[username]
	if !ok || CompareHashedPassword(password, user.HashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username/password"})
		return
	}

	//toke generation
	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

	// cookie
	c.SetCookie("session_token", sessionToken, 3600*24, "/", "", false, true)
	c.SetCookie("csrf_token", csrfToken, 3600*24, "/", "", false, true)

	//add into DB
	user.CSRFToken = csrfToken
	user.SessionToken = sessionToken
	usersDB[username] = user

	c.JSON(http.StatusOK, gin.H{"success": "login completed"})

}

// logout
func LogoutHandler(c *gin.Context) {

	// check auth
	if err := Authorize(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorised ",
		})
	}

	//clear cookie
	c.SetCookie("session_token", "", -1, "/", "", false, true)
	c.SetCookie("csrf_token", "", -1, "/", "", false, false)

	username := c.PostForm("username")
	user := usersDB[username]
	user.SessionToken = ""
	user.CSRFToken = ""
	usersDB[username] = user

	c.JSON(http.StatusOK, gin.H{"success": "logout completed"})

}

// Protected route
func HandleProtectedRoute(c *gin.Context) {
	if err := Authorize(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	username := c.PostForm("username")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Welcome user: %s", username)})
}

// home page

func ProductHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"messgae": "you are on product page"})
}
