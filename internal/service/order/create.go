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
