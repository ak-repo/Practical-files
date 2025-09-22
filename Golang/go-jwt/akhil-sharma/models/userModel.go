package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string `json:"username" validate:"required"`
	Password      string `json:"passwrod" validate:"required"`
	Token         string `json:"toke"`
	Refresh_token string `json:"refresh_token"`
	User_type     string `json:"user_type" validate:"required"`
	UID           string `json:"uid"`
}
