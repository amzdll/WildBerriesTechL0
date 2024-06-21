package order

import (
	"context"
	"encoding/json"
	"wb/internal/model"
	"wb/pkg/logger"

	"github.com/redis/go-redis/v9"
)

type Repository struct {
	logger *logger.Logger
	cache  *redis.Client
}

func New(logger *logger.Logger, cache *redis.Client) *Repository {
	repository := Repository{logger: logger, cache: cache}
	return &repository
}

func (r *Repository) LoadCache(ctx context.Context, orders []model.Order) {
	for _, order := range orders {
		data, err := json.Marshal(order)
		if err != nil {
			r.logger.Fatal("Can't load cache", err)

			err = r.cache.Set(ctx, order.OrderUID, data, 0).Err()

			if err != nil {
				r.logger.Fatal("Can't load cache", err)
			}
		}
	}
}
