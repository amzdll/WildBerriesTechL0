package entity

import "wb/internal/model"

type Payment struct {
	Transaction  string
	RequestId    string
	Currency     string
	Provider     string
	Amount       int
	PaymentDt    int64
	Bank         string
	DeliveryCost float64
	GoodsTotal   int
	CustomFee    float64
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
