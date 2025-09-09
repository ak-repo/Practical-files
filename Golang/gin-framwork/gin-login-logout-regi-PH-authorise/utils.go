package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// hash algorithum
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// compare hashed password
func CompareHashedPassword(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}

// generate random token
func generateToken(length int) string {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Token generate failed! %v", err)
	}

	return base64.URLEncoding.EncodeToString(bytes)
}
