package order

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"wb/pkg/logger"
)

type Repository struct {
	logger *logger.Logger
	db     *pgxpool.Pool
	cache  *redis.Client
}

func New(logger *logger.Logger, db *pgxpool.Pool, cache *redis.Client) *Repository {
	return &Repository{logger: logger, db: db, cache: cache}
}
