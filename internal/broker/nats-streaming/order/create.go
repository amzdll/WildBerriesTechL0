package order

import (
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"wb/internal/broker/nats-streaming/order/message"
)

func (h MessageHandler) create(m *stan.Msg) {
	var order message.Order
	if err := json.Unmarshal(m.Data, &order); err != nil {
		h.logger.Error("Can't unmarshall data", err)
		return
	}
	if err := h.validator.Struct(order); err != nil {
		h.logger.Error("Invalid data", err)
		return
	}
	if err := h.service.Create(context.Background(), order.ToModel()); err != nil {
		h.logger.Error("Data processing error", err)
		return
	}
	if err := m.Ack(); err != nil {
		h.logger.Error("Data processing confirmation error", err)
	}
}
