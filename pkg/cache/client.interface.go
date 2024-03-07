package cache

import (
	"time"
)

type Client interface {
	Set(key string, value any, expiration time.Duration) error
	Get(key string) (any, error)
	Incr(key string) error
}
