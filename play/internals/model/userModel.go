package model

import "time"

type UserTable struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email"`
	HashPassword string    `json:"-"`
}

type Input struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
