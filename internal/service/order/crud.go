package order

import (
	"context"
	"wb/internal/model"
)

func (s Service) Create(ctx context.Context, order model.Order) error {
	if err := s.db.Create(ctx, order); err != nil {
		return err
	}
	if err := s.cache.Create(ctx, order); err != nil {
		return err
	}
	return nil
}

func (s Service) Get(ctx context.Context, id string) (model.Order, error) {
	return s.cache.Get(ctx, id)
}
