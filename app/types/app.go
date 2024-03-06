package types

import "github.com/gofiber/fiber/v2"

type App struct {
	*fiber.App
	JWTSecret           string
	JWTExpiresInMinutes int
}
