package api

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/internal/entity"
	"rate-limiter/internal/usecase"
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
