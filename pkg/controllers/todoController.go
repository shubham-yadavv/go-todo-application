package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/shubham-yadavv/go-todo-application/pkg/config"
	"github.com/shubham-yadavv/go-todo-application/pkg/helpers"
	"github.com/shubham-yadavv/go-todo-application/pkg/models"
)

func GetAllTodosByUser(c *gin.Context) {
	userID := helpers.GetUserIDFromToken(c)
	completed := c.Query("completed")

	var todos []models.Todo

	switch completed {
	case "true":
		config.DB.Where("user_id = ? AND completed = ?", userID, true).Find(&todos)
	case "false":
		config.DB.Where("user_id = ? AND completed = ?", userID, false).Find(&todos)
	default:
		config.DB.Where("user_id = ?", userID).Find(&todos)
	}

	c.JSON(200, gin.H{
		"message": "Todos fetched successfully",
		"todos":   todos,
	})

}

func CreateTodo(c *gin.Context) {
	userID := helpers.GetUserIDFromToken(c)

	var body struct {
		Title     string
		Completed bool
		UserID    uint
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	todo := models.Todo{
		Title:     body.Title,
		Completed: body.Completed,
		UserID:    userID,
	}

	config.DB.Create(&todo)

	c.JSON(201, gin.H{
		"message": "Todo created successfully",
		"todo":    todo,
	})

}

func UpdateTodo(c *gin.Context) {
	var body struct {
		Title     string
		Completed bool
		UserID    uint
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	var todo models.Todo
	config.DB.First(&todo, c.Param("id"))

	if todo.ID == 0 {
		c.JSON(404, gin.H{
			"message": "Todo not found",
		})
		return
	}

	todo.Title = body.Title
	todo.Completed = body.Completed
	todo.UserID = body.UserID

	config.DB.Save(&todo)

	c.JSON(200, gin.H{
		"message": "Todo updated successfully",
		"todo":    todo,
	})

}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	config.DB.First(&todo, c.Param("id"))

	if todo.ID == 0 {
		c.JSON(404, gin.H{
			"message": "Todo not found",
		})
		return
	}

	config.DB.Delete(&todo)

	c.JSON(200, gin.H{
		"message": "Todo deleted successfully",
	})

}
