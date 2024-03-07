package main

import (
	"log"
	configuration "rate-limiter/config"
	"rate-limiter/internal/infra/rest"
	"rate-limiter/pkg/cache"
)

func main() {
	cfg, err := configuration.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	cacheClient := cache.NewRedisClient(cfg.RedisAddress, cfg.RedisPassword, 0)

	log.Fatalln(rest.CreateRestServer(cfg, cacheClient))
}
