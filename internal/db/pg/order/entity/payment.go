package entity

import "wb/internal/model"

type Payment struct {
	Transaction  string  `db:"transaction"`
	RequestId    string  `db:"request_id"`
	Currency     string  `db:"currency"`
	Provider     string  `db:"provider"`
	Amount       int     `db:"amount"`
	PaymentDt    int64   `db:"payment_dt"`
	Bank         string  `db:"bank"`
	DeliveryCost float64 `db:"delivery_cost"`
	GoodsTotal   int     `db:"goods_total"`
	CustomFee    float64 `db:"custom_fee"`
}

func (e *Payment) toModel() model.Payment {
	return model.Payment{
		Transaction:  e.Transaction,
		RequestId:    e.RequestId,
		Currency:     e.Currency,
		Provider:     e.Provider,
		Amount:       e.Amount,
		PaymentDt:    e.PaymentDt,
		Bank:         e.Bank,
		DeliveryCost: e.DeliveryCost,
		GoodsTotal:   e.GoodsTotal,
		CustomFee:    e.CustomFee,
	}
}
