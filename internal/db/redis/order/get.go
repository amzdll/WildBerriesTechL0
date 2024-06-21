package order

import (
	"context"
	"encoding/json"
	"fmt"
	"wb/internal/db/pg"
	"wb/internal/model"
)

func (r *Repository) Get(ctx context.Context, id string) (model.Order, error) {
	var order model.Order

	data, err := r.cache.Get(ctx, id).Result()
	if err != nil {
		return model.Order{}, fmt.Errorf("get order: %w", pg.ErrNotFound)
	}

	if err = json.Unmarshal([]byte(data), &order); err != nil {
		return model.Order{}, err
	}

	return order, nil
}
