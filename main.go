package main

import (
	"log"
	configuration "rate-limiter/config"
	"rate-limiter/internal/infra/rest"
	"rate-limiter/pkg/cache"
)

func main() {
	config, err := configuration.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	cacheClient := cache.NewRedisClient(config.RedisAddress, config.RedisPassword, 0)

	log.Fatalln(rest.CreateRestServer(config, cacheClient))
}
