package main

import (
	"errors"
	"net/url"

	"github.com/gin-gonic/gin"
)

// authetication middleware
var AuthError = errors.New("Unauthorised")

func Authorize(c *gin.Context) error {

	username := c.PostForm("username")
	user, ok := usersDB[username]
	if !ok {
		return AuthError
	}

	//checkn session cookie
	st, err := c.Cookie("session_token")
	if err != nil || st == "" || st != user.SessionToken {
		return AuthError
	}

	//check CSRF header
	csrf := c.GetHeader("X-CSRF-Token")
	decodedCSRF, _ := url.QueryUnescape(csrf)
	if decodedCSRF != user.CSRFToken || csrf == "" {
		return AuthError
	}

	return nil
}
