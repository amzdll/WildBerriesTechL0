package message

import "wb/internal/model"

type Payment struct {
	Transaction  string  `json:"transaction"  validate:"required"`
	RequestId    string  `json:"request_id"  validate:"required"`
	Currency     string  `json:"currency"  validate:"required"`
	Provider     string  `json:"provider"  validate:"required"`
	Amount       int     `json:"amount"  validate:"required"`
	PaymentDt    int64   `json:"payment_dt"  validate:"required"`
	Bank         string  `json:"bank"  validate:"required"`
	DeliveryCost float64 `json:"delivery_cost"  validate:"required"`
	GoodsTotal   int     `json:"goods_total"  validate:"required"`
	CustomFee    float64 `json:"custom_fee"  validate:"required"`
}

func (p *Payment) toModel() model.Payment {
	return model.Payment{
		Transaction:  p.Transaction,
		RequestId:    p.RequestId,
		Currency:     p.Currency,
		Provider:     p.Provider,
		Amount:       p.Amount,
		PaymentDt:    p.PaymentDt,
		Bank:         p.Bank,
		DeliveryCost: p.DeliveryCost,
		GoodsTotal:   p.GoodsTotal,
		CustomFee:    p.CustomFee,
	}
}
