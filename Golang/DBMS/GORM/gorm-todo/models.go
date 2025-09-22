package main

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Content string `json:"content"`
	Done bool `json:"done"`
}
