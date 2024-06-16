package order

import (
	"context"
	"encoding/json"
	"fmt"
	"wb/internal/db/pg"
	"wb/internal/model"
)

func (r *Repository) Create(ctx context.Context, order model.Order) error {
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}
	if err = r.cache.Set(ctx, order.OrderUID, orderJSON, 0).Err(); err != nil {
		return fmt.Errorf("get order: %w", pg.ErrDuplicateKey)
	}

	return nil
}

func (r *Repository) CreateBatch(ctx context.Context, orders []model.Order) {
	for _, order := range orders {
		if err := r.Create(ctx, order); err != nil {
			r.logger.Error("Failed to create record", err)
			continue
		}
	}
}
