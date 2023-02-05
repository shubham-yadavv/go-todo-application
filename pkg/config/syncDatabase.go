package config

import "github.com/shubham-yadavv/go-todo-application/pkg/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Todo{})
	DB.AutoMigrate(&models.User{})
}
