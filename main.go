package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"rate-limiter/app/types"
	configuration "rate-limiter/config"
	"rate-limiter/routes"
)

func main() {
	config, err := configuration.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	app := types.App{
		App:                 fiber.New(),
		JWTSecret:           config.JWTSecret,
		JWTExpiresInMinutes: config.JWTExpiresInMinutes,
	}

	app.Use(logger.New())
	app.Use(requestid.New())

	api := app.Group("/api")
	apiV1 := api.Group("/v1")
	routes.RegisterAPI(apiV1, &app)

	log.Fatalln(app.Listen(config.WebServerPort))
}
