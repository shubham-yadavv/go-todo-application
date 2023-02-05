package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham-yadavv/go-todo-application/pkg/controllers"
	"github.com/shubham-yadavv/go-todo-application/pkg/middleware"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/signup", controllers.SignUp)
	incommingRoutes.POST("/login", controllers.Login)

	incommingRoutes.GET("/me", middleware.Authenticate, controllers.Profile)

}
