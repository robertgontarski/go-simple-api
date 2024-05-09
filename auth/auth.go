package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"simple-api/models"
	"time"
)

type JWTAuth struct {
	private string
}

func NewJWTAuth() *JWTAuth {
	return &JWTAuth{
		private: os.Getenv("JWT_PRIVATE_KEY"),
	}
}

func (a *JWTAuth) CreateToken(client *models.Client) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"id":  client.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return t.SignedString([]byte(a.private))
}

func (a *JWTAuth) VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(a.private), nil
	})
}
