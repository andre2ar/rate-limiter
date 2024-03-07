package entity

import (
	"github.com/andre2ar/rate-limiter/pkg/cache"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	*fiber.App
	CacheClient         cache.Client
	JWTSecret           string
	JWTExpiresInMinutes int
}
