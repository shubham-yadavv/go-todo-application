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

	routes.UserRoutes(r)
	routes.NotesRoutes(r)

	r.GET("/", heatlhCheck)

	r.Run("localhost:3000")

}

func heatlhCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
