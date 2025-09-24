package models

import (
	"errors"
	"regexp"

	"gorm.io/gorm"
)

// User represents in DB
type User struct {
	gorm.Model
	Email        string `json:"email" gorm:"uniqueIndex;size:255;not null"`
	PasswordHash string `json:"-" gorm:"not null"`
	Role         string `json:"role" gorm:"not null"`
}

// InputUser represents login/registration payload
type InputUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Response carry
type Response struct {
	RefreshToken string
	AccessToken  string
	User         *User
}

// Validate checks input fields
func (i *InputUser) Validate() error {
	// More flexible email regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(i.Email) {
		return errors.New("invalid email format")
	}

	// Password validation
	if len(i.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}
