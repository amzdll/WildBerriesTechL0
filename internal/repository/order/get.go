package order

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"wb/internal/model"
)

func (r *Repository) Get(ctx context.Context, id string) (model.Order, error) {
	var order model.Order
	var err error

	if err = pgxscan.Get(ctx, r.db, &order, getOrderQuery, id); err != nil {
		return model.Order{}, err
	}
	if err = pgxscan.Get(ctx, r.db, &order.Delivery, getDeliveryQuery, id); err != nil {
		return model.Order{}, err
	}
	if err = pgxscan.Get(ctx, r.db, &order.Payment, getPaymentQuery, id); err != nil {
		return model.Order{}, err
	}
	if err = pgxscan.Select(ctx, r.db, &order.Items, getItemQuery, order.TrackNumber); err != nil {
		fmt.Println(err)
		return model.Order{}, err
	}
	fmt.Println(order.Items)
	return order, nil
}
