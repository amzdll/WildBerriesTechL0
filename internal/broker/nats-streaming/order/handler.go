package order

import (
	"context"
	"github.com/nats-io/stan.go"
	"time"
	"wb/internal/model"
	"wb/pkg/logger"
)

type Service interface {
	Create(ctx context.Context, order model.Order) error
}

type MessageHandler struct {
	service Service
	logger  *logger.Logger
}

func New(service Service, logger *logger.Logger) *MessageHandler {
	return &MessageHandler{
		service: service,
		logger:  logger,
	}
}

func (h MessageHandler) Subscribe(conn stan.Conn) error {
	var err error
	if _, err = conn.Subscribe(
		"orders",
		h.Create,
		stan.DeliverAllAvailable(),
		stan.AckWait(30*time.Second),
		stan.SetManualAckMode(),
	); err != nil {
		h.logger.Fatal("Error when connecting to a channel", err)
	}
	return err
}