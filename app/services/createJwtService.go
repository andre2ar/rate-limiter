package services

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateJWT(secret string, expiresInMinutes int) (string, error) {
	claims := jwt.MapClaims{
		"name": "John Doe",
		"exp":  time.Now().Add(time.Minute * time.Duration(expiresInMinutes)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	return t, err
}
