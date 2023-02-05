package helpers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserIDFromToken(c *gin.Context) uint {
	tokenString, err := c.Cookie("token")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return 0
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return uint(claims["user_id"].(float64))
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return 0
	}

}
