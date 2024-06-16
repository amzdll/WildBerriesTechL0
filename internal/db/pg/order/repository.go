package order

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"wb/pkg/logger"
)

type Repository struct {
	logger *logger.Logger
	db     *pgxpool.Pool
}

func New(logger *logger.Logger, db *pgxpool.Pool) *Repository {
	repository := Repository{logger: logger, db: db}
	return &repository
}
