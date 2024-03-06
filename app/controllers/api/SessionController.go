package api

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/app/services"
	"rate-limiter/app/types"
)

func CreateSession(app *types.App) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jwt, err := services.CreateJWT(app.JWTSecret, app.JWTExpiresInMinutes)
		if err != nil {
			return err
		}

		err = ctx.JSON(fiber.Map{
			"access_token": jwt,
		})

		return err
	}
}
