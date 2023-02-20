package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham-yadavv/go-todo-application/pkg/controllers"
	"github.com/shubham-yadavv/go-todo-application/pkg/middleware"
)

func NotesRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/notes", middleware.Authenticate, controllers.GetNotes)
	incommingRoutes.POST("/notes", middleware.Authenticate, controllers.CreateNotes)

}
