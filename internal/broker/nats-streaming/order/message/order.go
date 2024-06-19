package message

import (
	"time"
	"wb/internal/model"
)

type Order struct {
	OrderUID          string    `json:"order_uid"  validate:"required"`
	TrackNumber       string    `json:"track_number"  validate:"required"`
	Entry             string    `json:"entry"  validate:"required"`
	Delivery          Delivery  `json:"delivery"  validate:"required"`
	Payment           Payment   `json:"payment"  validate:"required"`
	Items             []Item    `json:"items"  validate:"required"`
	Locale            string    `json:"locale"  validate:"required"`
	InternalSignature string    `json:"internal_signature"  validate:"required"`
	CustomerID        string    `json:"customer_id"  validate:"required"`
	DeliveryService   string    `json:"delivery_service"  validate:"required"`
	ShardKey          string    `json:"shardkey"  validate:"required"`
	SmID              int       `json:"sm_id"  validate:"required"`
	DateCreated       time.Time `json:"date_created"  validate:"required"`
	OofShard          string    `json:"oof_shard"  validate:"required"`
}

func (o *Order) ToModel() model.Order {
	items := make([]model.Item, len(o.Items))
	for i, itemOrder := range o.Items {
		items[i] = itemOrder.toModel()
	}
	return model.Order{
		OrderUID:          o.OrderUID,
		TrackNumber:       o.TrackNumber,
		Entry:             o.Entry,
		Delivery:          o.Delivery.toModel(),
		Payment:           o.Payment.toModel(),
		Items:             items,
		Locale:            o.Locale,
		InternalSignature: o.InternalSignature,
		CustomerID:        o.CustomerID,
		DeliveryService:   o.DeliveryService,
		ShardKey:          o.ShardKey,
		SmID:              o.SmID,
		DateCreated:       o.DateCreated,
		OofShard:          o.OofShard,
	}
}
