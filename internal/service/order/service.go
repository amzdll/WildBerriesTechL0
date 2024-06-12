package order

import (
	"context"
	"wb/internal/model"
)

type Repository interface {
	Create(ctx context.Context, order model.Order) error
	Get(ctx context.Context, id string) (model.Order, error)
}

type Service struct {
	repository Repository
}

func New(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
