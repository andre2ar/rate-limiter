package main

import (
	configuration "github.com/andre2ar/rate-limiter/config"
	"github.com/andre2ar/rate-limiter/internal/infra/rest"
	"github.com/andre2ar/rate-limiter/pkg/cache"
	"log"
)

func main() {
	cfg, err := configuration.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	cacheClient := cache.NewRedisClient(cfg.RedisAddress, cfg.RedisPassword, 0)

	log.Fatalln(rest.CreateRestServer(cfg, cacheClient))
}
