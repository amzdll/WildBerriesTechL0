package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"wb/internal/config"
	"wb/pkg/logger"
)

func redisModule() fx.Option {
	return fx.Module(
		"Redis",
		fx.Provide(
			connect,
			config.NewRedisConfig,
		),
		fx.Invoke(closeConnection),
	)
}

func connect(config *config.RedisConfig, log *logger.Logger) *redis.Client {
	ctx := context.Background()
	drb := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})
	if err := drb.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis.", err)
	}
	log.Info("Redis connection has opened.")
	return drb
}

func closeConnection(lc fx.Lifecycle, connection *redis.Client, log *logger.Logger) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			defer func(connection *redis.Client) {
				err := connection.Close()
				if err != nil {
					log.Error("Redis connection close error:", err)
				}
			}(connection)
			log.Info("Redis connection closed.")
			return nil
		},
	})
}
