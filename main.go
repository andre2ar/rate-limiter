package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	configuration "rate-limiter/config"
	"rate-limiter/routes"
)

func main() {
	config, err := configuration.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	app := fiber.New()

	app.Use(logger.New())
	app.Use(requestid.New())

	api := app.Group("/api")
	apiV1 := api.Group("/v1")
	routes.RegisterAPI(apiV1)

	log.Fatalln(app.Listen(config.WebServerPort))
}
