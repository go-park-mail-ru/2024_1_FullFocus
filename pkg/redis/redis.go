package redis

import (
	"github.com/go-redis/redis"

	"github.com/go-park-mail-ru/2024_1_FullFocus/internal/config"
)

func NewClient(cfg config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Port,
		Password: "",
		DB:       0,
	})
}
