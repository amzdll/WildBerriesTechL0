package entity

import (
	"time"
	"wb/internal/model"
)

type Order struct {
	OrderUID          string `db:"order_uid"`
	TrackNumber       string `db:"track_number"`
	Entry             string `db:"entry"`
	Delivery          Delivery
	Payment           Payment
	Items             []Item
	Locale            string    `db:"locale"`
	InternalSignature string    `db:"internal_signature"`
	CustomerID        string    `db:"customer_id"`
	DeliveryService   string    `db:"delivery_service"`
	ShardKey          string    `db:"shardkey"`
	SmID              int       `db:"sm_id"`
	DateCreated       time.Time `db:"date_created"`
	OofShard          string    `db:"oof_shard"`
}

func (e *Order) ToModel() model.Order {
	items := make([]model.Item, len(e.Items))
	for i, item := range e.Items {
		items[i] = item.toModel()
	}

	return model.Order{
		OrderUID:          e.OrderUID,
		TrackNumber:       e.TrackNumber,
		Entry:             e.Entry,
		Delivery:          e.Delivery.toModel(),
		Payment:           e.Payment.toModel(),
		Items:             items,
		Locale:            e.Locale,
		InternalSignature: e.InternalSignature,
		CustomerID:        e.CustomerID,
		DeliveryService:   e.DeliveryService,
		ShardKey:          e.ShardKey,
		SmID:              e.SmID,
		DateCreated:       e.DateCreated,
		OofShard:          e.OofShard,
	}
}
