package order

import (
	"context"
	"wb/internal/model"
)

func (s Service) Get(ctx context.Context, id string) (model.Order, error) {
	return s.cache.Get(ctx, id)
}
