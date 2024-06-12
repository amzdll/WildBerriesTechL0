package order

import (
	"context"
	"wb/internal/model"
)

func (r *Repository) Create(ctx context.Context, order model.Order) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
			return
		}
		err = tx.Commit(ctx)
	}()

	if err = r.createOrder(ctx, order); err != nil {
		return err
	}
	if err = r.createDelivery(ctx, order.Delivery); err != nil {
		return err
	}
	if err = r.createItems(ctx, order); err != nil {
		return err
	}
	if err = r.createPayment(ctx, order.Payment); err != nil {
		return err
	}

	return nil
}

func (r *Repository) createOrder(ctx context.Context, order model.Order) error {
	_, err := r.db.Exec(
		ctx, createOrderQuery,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.ShardKey,
		order.SmID,
		order.DateCreated,
		order.OofShard,
	)
	return err
}

func (r *Repository) createDelivery(ctx context.Context, delivery model.Delivery) error {
	_, err := r.db.Exec(
		ctx, createDeliveryQuery,
		delivery.OrderUID,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email,
	)
	return err
}

func (r *Repository) createItems(ctx context.Context, order model.Order) error {
	var err error
	for _, item := range order.Items {
		_, err = r.db.Exec(
			ctx, createItemQuery,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.RID,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			item.Brand,
			item.Status)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) createPayment(ctx context.Context, payment model.Payment) error {
	_, err := r.db.Exec(
		ctx, createPaymentQuery,
		payment.Transaction,
		payment.RequestId,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	)
	return err
}
