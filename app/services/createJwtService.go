package services

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateJWT() (string, error) {
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("very_public_secret"))

	return t, err
}
