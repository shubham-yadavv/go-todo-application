package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shubham-yadavv/go-todo-application/pkg/config"
	"github.com/shubham-yadavv/go-todo-application/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	var existingUser models.User

	if config.DB.Where("username = ?", body.Username).First(&existingUser).RowsAffected != 0 {
		c.JSON(400, gin.H{
			"message": "Username already exists",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 8)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := models.User{
		Username: body.Username,
		Password: string(hashedPassword),
	}

	config.DB.Create(&user)

	c.JSON(201, gin.H{
		"message": "User created successfully",
		"user":    user,
	})

}

func Login(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	var user models.User

	if config.DB.Where("username = ?", body.Username).First(&user).RowsAffected == 0 {
		c.JSON(400, gin.H{
			"message": "Username does not exist",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.UserID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.SetCookie("token", tokenString, 60*60*24*30, "/", "localhost", false, true)

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Login successful",
		tokenString: tokenString,
	})

}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{
		"message": "Logout successful",
	})
}

func GetAuthUser(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var user models.User
	config.DB.Where("user_id = ?", userID).First(&user)
	c.JSON(200, gin.H{
		"user": user,
	})
}
