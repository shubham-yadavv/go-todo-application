package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserID    uint   `json:"user_id"`
}
