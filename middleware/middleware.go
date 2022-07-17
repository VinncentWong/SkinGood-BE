package middleware

import (
	"module/domain"
	"os"
	"time"

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
