package order

import (
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"wb/internal/broker/nats-streaming/order/message"
)

func (h MessageHandler) Create(m *stan.Msg) {
	var order message.Order
	if err := json.Unmarshal(m.Data, &order); err != nil {
		return
	}
	if err := h.service.Create(context.Background(), order.ToModel()); err != nil {
		log.Println(order.OrderUID)
		h.logger.Error("Data processing error", err)
		return
	}

	if err := m.Ack(); err != nil {
		h.logger.Error("Data processing confirmation error", err)
	}
	h.logger.Info("Order created")
}
