package routes

import (

	// gin
	"github.com/gin-gonic/gin"
	"github.com/shubham-yadavv/go-todo-application/pkg/controllers"
	"github.com/shubham-yadavv/go-todo-application/pkg/middleware"
)

func TodoRoutes(incommingRoutes *gin.Engine) {

	// get all todos of a user
	incommingRoutes.GET("/todos", middleware.Authenticate, controllers.GetAllTodosByUser)

	incommingRoutes.POST("/todos", controllers.CreateTodo)
	incommingRoutes.PUT("/todos", controllers.UpdateTodo)
	incommingRoutes.DELETE("/todos", controllers.DeleteTodo)
	incommingRoutes.GET("/me/:id", middleware.Authenticate, controllers.Profile)

}
