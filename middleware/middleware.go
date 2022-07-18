package middleware

import (
	"module/domain"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user domain.User) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.MembershipID,
		"created-at": user.CreatedAt,
		"gender":     user.CreatedAt,
		"exp":        time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	finalToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &finalToken, nil
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		token := bearer[7:]
		fixedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  err.Error(),
				"success": false,
			})
			return
		}
		_, ok := fixedToken.Claims.(jwt.MapClaims)
		if fixedToken.Valid && ok {
			c.Next()
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "token doesn't valid",
				"success": false,
			})
			c.Abort()
			return
		}
	}
}
