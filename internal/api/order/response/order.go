package response

import (
	"time"
	"wb/internal/model"
)

type Order struct {
	OrderUID          string
	TrackNumber       string
	Entry             string
	Delivery          delivery
	Payment           payment
	Items             []item
	Locale            string
	InternalSignature string
	CustomerID        string
	DeliveryService   string
	Shardkey          string
	SmID              int
	DateCreated       time.Time
	OofShard          string
}

func OrderFromModel(m model.Order) Order {
	res := Order{
		OrderUID:          m.OrderUID,
		TrackNumber:       m.TrackNumber,
		Entry:             m.Entry,
		Locale:            m.Locale,
		InternalSignature: m.InternalSignature,
		CustomerID:        m.CustomerID,
		DeliveryService:   m.DeliveryService,
		Shardkey:          m.ShardKey,
		SmID:              m.SmID,
		DateCreated:       m.DateCreated,
		OofShard:          m.OofShard,
	}
	res.Delivery.fromModel(m.Delivery)
	res.Payment.fromModel(m.Payment)

	for _, mItem := range m.Items {
		var rItem item

		rItem.fromModel(mItem)
		res.Items = append(res.Items, rItem)
	}

	return res
}
