package routes

import (
	"github.com/andre2ar/rate-limiter/internal/entity"
	"github.com/andre2ar/rate-limiter/internal/infra/rest/controllers/api"
	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(router fiber.Router, app *entity.App) {
	router.Post("/session", api.CreateSession(app))
}
