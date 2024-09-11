package test

import (
	"context"
	"github.com/nopecho/golang-template/test/testdata"
	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

const (
	redisImage = "redis:7-alpine"
)

type RedisContainer struct {
	Container *redis.RedisContainer
	Context   context.Context
	Endpoint  string
}

func NewRedisContainer() *RedisContainer {
	ctx := context.Background()
	redisContainer, err := redis.Run(ctx,
		redisImage,
		redis.WithLogLevel(redis.LogLevelVerbose),
		redis.WithConfigFile(testdata.Redis),
	)
	if err != nil {
		log.Fatal().Msgf("failed to start redis container: %s", err)
	}

	endpoint, err := redisContainer.Endpoint(ctx, "")
	if err != nil {
		log.Fatal().Msgf("failed to get redis container endpoint: %s", err)
	}
	return &RedisContainer{
		Container: redisContainer,
		Context:   ctx,
		Endpoint:  endpoint,
	}
}

func (r *RedisContainer) Terminate() {
	if err := r.Container.Terminate(r.Context); err != nil {
		log.Fatal().Msgf("failed to terminate redis container: %s", err)
	}
}
