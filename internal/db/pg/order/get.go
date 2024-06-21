package order

import (
	"context"
	"wb/internal/db/pg/order/entity"
	"wb/internal/db/pg/order/query"
	"wb/internal/model"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *Repository) GetAll(ctx context.Context) ([]model.Order, error) {
	var err error

	orders := make([]entity.Order, 0)
	if err = pgxscan.Select(ctx, r.db, &orders, query.GetOrders); err != nil {
		return []model.Order{}, err
	}

	for i := range orders {
		if err = pgxscan.Get(ctx, r.db, &orders[i].Delivery, query.GetDeliveryById, orders[i].OrderUID); err != nil {
			orders[i] = entity.Order{}
			continue
		}

		if err = pgxscan.Get(ctx, r.db, &orders[i].Payment, query.GetPaymentById, orders[i].OrderUID); err != nil {
			orders[i] = entity.Order{}
			continue
		}

		if err = pgxscan.Select(ctx, r.db, &orders[i].Items, query.GetItemById, orders[i].TrackNumber); err != nil {
			orders[i] = entity.Order{}
			continue
		}
	}

	result := make([]model.Order, len(orders))
	for i, o := range orders {
		result[i] = o.ToModel()
	}

	return result, nil
}
