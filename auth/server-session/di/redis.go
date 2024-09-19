package di

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type RedisConfig struct {
	RedisUrl string
}

func newRedisClient(cfg RedisConfig) *redis.Client {
	opts, err := redis.ParseURL(cfg.RedisUrl)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse redis url")
	}

	rdb := redis.NewClient(opts)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal().Err(err).Msg("failed to ping redis")
	}

	return rdb
}
