package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"rate-limiter/config"
	"rate-limiter/internal/entity"
	"rate-limiter/internal/infra/rest/middleware/ratelimiter"
	"rate-limiter/internal/infra/rest/routes"
	"rate-limiter/pkg/cache"
)

func CreateRestServer(cfg *config.Config, cacheClient cache.Client) error {
	app := entity.App{
		App:                 fiber.New(),
		CacheClient:         cacheClient,
		JWTSecret:           cfg.JWTSecret,
		JWTExpiresInMinutes: cfg.JWTExpiresInMinutes,
	}

	app.Use(logger.New())
	app.Use(requestid.New())
	app.Use(ratelimiter.New(app.CacheClient, cfg.MaxRequestPerMinuteIP, cfg.MaxRequestPerMinuteToken))

	api := app.Group("/api")
	apiV1 := api.Group("/v1")
	routes.RegisterAPI(apiV1, &app)

	err := app.Listen(cfg.WebServerPort)
	if err != nil {
		return err
	}

	return nil
}
