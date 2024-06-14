package response

import "wb/internal/model"

type payment struct {
	Transaction  string  `json:"transaction"`
	RequestId    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       int     `json:"amount"`
	PaymentDt    int64   `json:"payment_dt"`
	Bank         string  `json:"bank"`
	DeliveryCost float64 `json:"delivery_cost"`
	GoodsTotal   int     `json:"goods_total"`
	CustomFee    float64 `json:"custom_fee"`
}

func (p *payment) fromModel(m model.Payment) {
	p.Transaction = m.Transaction
	p.RequestId = m.RequestId
	p.Currency = m.Currency
	p.Provider = m.Provider
	p.Amount = m.Amount
	p.PaymentDt = m.PaymentDt
	p.Bank = m.Bank
	p.DeliveryCost = m.DeliveryCost
	p.GoodsTotal = m.GoodsTotal
	p.CustomFee = m.CustomFee
}
