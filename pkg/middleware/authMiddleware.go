package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shubham-yadavv/go-todo-application/pkg/config"
	"github.com/shubham-yadavv/go-todo-application/pkg/models"
)

func Authenticate(c *gin.Context) {
	tokenString, err := c.Cookie("token")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		if config.DB.Where("username = ?", claims["username"]).First(&user).RowsAffected == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
