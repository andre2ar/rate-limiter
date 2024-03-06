package routes

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/app/controllers/api"
)

func RegisterAPI(router fiber.Router) {
	router.Post("/session", api.CreateSession())
}
