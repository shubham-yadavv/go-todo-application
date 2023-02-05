package main

import (
	"github.com/gin-gonic/gin"

	"github.com/shubham-yadavv/go-todo-application/pkg/config"
	"github.com/shubham-yadavv/go-todo-application/pkg/routes"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
	config.SyncDatabase()
}

func main() {
	r := gin.Default()

	routes.TodoRoutes(r)
	routes.UserRoutes(r)

	r.Run("localhost:3000")

}