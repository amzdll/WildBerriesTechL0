package order

import (
	"context"
	"wb/internal/model"
	"wb/pkg/logger"
)

type PgRepository interface {
	Create(ctx context.Context, order model.Order) error
	GetAll(ctx context.Context) ([]model.Order, error)
}

type RedisRepository interface {
	Create(ctx context.Context, order model.Order) error
	CreateBatch(ctx context.Context, orders []model.Order)
	Get(ctx context.Context, id string) (model.Order, error)
}

type Service struct {
	logger *logger.Logger
	db     PgRepository
	cache  RedisRepository
}

func New(logger *logger.Logger, cacheRepository RedisRepository, pgRepository PgRepository) *Service {
	s := Service{
		logger: logger,
		cache:  cacheRepository,
		db:     pgRepository,
	}
	s.loadCache()
	return &s
}

func (s Service) loadCache() {
	ctx := context.Background()
	orders, err := s.db.GetAll(ctx)
	if err != nil {
		s.logger.Error("Error loading cache", err)
	}
	s.cache.CreateBatch(ctx, orders)
}
