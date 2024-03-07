package api

import (
	"github.com/andre2ar/rate-limiter/internal/entity"
	"github.com/andre2ar/rate-limiter/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

func CreateSession(app *entity.App) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jwt, err := usecase.CreateJWT(app.JWTSecret, app.JWTExpiresInMinutes)
		if err != nil {
			return err
		}

		err = ctx.JSON(fiber.Map{
			"access_token": jwt,
		})

		return err
	}
}
