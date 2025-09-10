package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// id generator function
func GenerateID() string {
	return uuid.New().String()
}

// password hasing algorithum
func GenerateHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// compare hashed password
func CompareHashAndPassword(password, hased string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hased))
	return err == nil
}


// token generation
func GenerateToken(length int) string {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Token generate Failed: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes)

}
