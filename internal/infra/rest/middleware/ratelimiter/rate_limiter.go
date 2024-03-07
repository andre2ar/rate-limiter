package ratelimiter

import (
	"github.com/gofiber/fiber/v2"
	"rate-limiter/pkg/cache"
	"strconv"
	"time"
)

type Data struct {
	IP           string
	RequestCount int
}

func New(cacheClient cache.Client, maxRequestPerMinuteIP int, maxRequestPerMinuteToken int) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		key := ctx.IP()
		maxRequestPerMinute := maxRequestPerMinuteIP

		apiKey := ctx.Get("API_KEY")
		if apiKey != "" {
			key = apiKey
			maxRequestPerMinute = maxRequestPerMinuteToken
		}

		value, err := cacheClient.Get(key)
		if err != nil {
			err := cacheClient.Set(key, 1, 1*time.Minute)
			if err != nil {
				return err
			}

			return ctx.Next()
		}

		accessCount, _ := strconv.Atoi(value.(string))
		if accessCount >= maxRequestPerMinute {
			return ctx.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "you have reached the maximum number of requests or actions allowed within a certain time frame",
			})
		}

		err = cacheClient.Incr(key)
		if err != nil {
			return err
		}

		return ctx.Next()
	}
}
