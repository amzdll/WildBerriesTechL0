package message

import (
	"time"
	"wb/internal/model"
)

type Order struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
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
