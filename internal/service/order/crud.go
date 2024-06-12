package order

import (
	"context"
	"wb/internal/model"
)

func (s Service) Create(ctx context.Context, order model.Order) error {
	return s.repository.Create(ctx, order)
}

func (s Service) Get(ctx context.Context, id string) (model.Order, error) {
	return s.repository.Get(ctx, id)
}
