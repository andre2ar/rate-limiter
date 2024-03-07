package routes

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/internal/entity"
	"rate-limiter/internal/infra/rest/controllers/api"
)

func RegisterAPI(router fiber.Router, app *entity.App) {
	router.Post("/session", api.CreateSession(app))
}
