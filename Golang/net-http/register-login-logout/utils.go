package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/crypto/bcrypt"
)

// password into hash
func hashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	log.Println(string(bytes))
	return string(bytes), err

}

// compare login password and hashed password
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// session token generation
func generateToken(length int) string {

	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}

	return base64.URLEncoding.EncodeToString(bytes)
}

// authentication handling
var AuthError = errors.New("Unauthorized")

func Authorize(r *http.Request) error {
	username := r.FormValue("username")
	user, ok := users[username]
	if !ok {
		return AuthError
	}

	//get the Session Token from the cookie
	st, err := r.Cookie("session_token")
	log.Println(st.Value == user.SessionToken)
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		return AuthError
	}

	// Get the CSRF token from the header
	csrf := r.Header.Get("X-CSRF-Token")
	//decode csrf
	decodedCSRF, _ := url.QueryUnescape(csrf)

	log.Println(user.CSRFToken)
	log.Println(decodedCSRF)

	if decodedCSRF != user.CSRFToken || csrf == "" {
		return AuthError
	}

	return nil

}
