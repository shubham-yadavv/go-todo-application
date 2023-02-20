package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham-yadavv/go-todo-application/pkg/config"
	"github.com/shubham-yadavv/go-todo-application/pkg/models"
)

func CreateNotes(c *gin.Context) {
	// get user id from token
	userID := c.MustGet("user_id").(uint)

	// get title and content from request body
	var body struct {
		Title   string
		Content string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// create note
	note := models.Notes{
		Title:   body.Title,
		Content: body.Content,
		UserID:  userID,
	}

	config.DB.Create(&note)

	c.JSON(201, gin.H{
		"message": "Note created successfully",
		"note":    note,
	})
}

func GetNotes(c *gin.Context) {
	// get user id from token
	userID := c.MustGet("user_id").(uint)

	// get all notes
	var notes []models.Notes
	config.DB.Where("user_id = ?", userID).Find(&notes)

	c.JSON(200, gin.H{
		"message": "Notes fetched successfully",
		"notes":   notes,
	})
}
