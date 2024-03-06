package routes

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/app/controllers/api"
	"rate-limiter/app/types"
)

func RegisterAPI(router fiber.Router, app *types.App) {
	router.Post("/session", api.CreateSession(app))
}
