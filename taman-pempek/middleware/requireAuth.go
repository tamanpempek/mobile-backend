package middleware

import (
	"fmt"
	"net/http"
	"os"
	"taman-pempek/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type middleware struct {
	userService user.UserService
}

func NewMiddleware(userService user.UserService) *middleware {
	return &middleware{userService}
}

func (m *middleware) RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Login first!",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"data":  nil,
				"msg":   "Login first!",
			})
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"data":  nil,
				"msg":   "Login first!",
			})
		}
		user, err := m.userService.FindUserByID(claims["foo"])

		if user.ID == 0 || err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"data":  nil,
				"msg":   "Login first!",
			})
		}

		c.Set("UserID", user.ID)
		c.Set("UserName", user.Name)
		c.Set("UserEmail", user.Email)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"data":  nil,
			"msg":   "Login first!",
		})
	}
}
