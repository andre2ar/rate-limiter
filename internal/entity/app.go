package entity

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/pkg/cache"
)

type App struct {
	*fiber.App
	CacheClient         cache.Client
	JWTSecret           string
	JWTExpiresInMinutes int
}
