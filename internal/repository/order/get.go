package order

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"wb/internal/model"
	"wb/internal/repository/order/entity"
)

func (r *Repository) Get(ctx context.Context, id string) (model.Order, error) {
	var order entity.Order
	if err := pgxscan.Get(ctx, r.db, &order, getOrderQuery, id); err != nil {
		return model.Order{}, err
	}
	if err := pgxscan.Get(ctx, r.db, &order.Delivery, getDeliveryQuery, id); err != nil {
		return model.Order{}, err
	}
	if err := pgxscan.Get(ctx, r.db, &order.Payment, getPaymentQuery, id); err != nil {
		return model.Order{}, err
	}
	if err := pgxscan.Select(ctx, r.db, &order.Items, getItemQuery, order.TrackNumber); err != nil {
		return model.Order{}, err
	}

	return order.ToModel(), nil
}
