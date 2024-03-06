package api

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/app/services"
)

func CreateSession() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		jwt, err := services.CreateJWT()
		if err != nil {
			return err
		}

		err = ctx.JSON(fiber.Map{
			"access_token": jwt,
		})

		return err
	}
}
