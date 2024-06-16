package order

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nats-io/stan.go"
	"wb/internal/broker/nats-streaming/order/message"
	"wb/internal/db/pg"
)

func (h MessageHandler) create(m *stan.Msg) {
	var order message.Order
	if err := json.Unmarshal(m.Data, &order); err != nil {
		h.logger.Error("Can't unmarshall data", err)
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

func (h MessageHandler) processPending(m *stan.Msg) {
	var order message.Order
	if err := json.Unmarshal(m.Data, &order); err != nil {
		h.logger.Error("Can't unmarshall data", err)
		return
	}
	err := h.service.Create(context.Background(), order.ToModel())
	if err != nil && !errors.Is(err, pg.ErrDuplicateKey) {
		h.logger.Error("Data processing error", err)
		return
	}

	if err = m.Ack(); err != nil {
		h.logger.Error("Data processing confirmation error", err)
	}
	h.logger.Info("Order created")
}
