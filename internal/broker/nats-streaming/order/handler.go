package order

import (
	"context"
	"time"
	"wb/internal/model"
	"wb/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
)

type Service interface {
	Create(ctx context.Context, order model.Order) error
}

type MessageHandler struct {
	service   Service
	validator *validator.Validate
	logger    *logger.Logger
}

func New(
	service Service,
	validator *validator.Validate,
	logger *logger.Logger,
) *MessageHandler {
	return &MessageHandler{
		service:   service,
		validator: validator,
		logger:    logger,
	}
}

func (h MessageHandler) Subscribe(conn stan.Conn) error {
	var err error

	const waitTime = 30 * time.Second
	if _, err = conn.Subscribe(
		"orders",
		h.create,
		stan.AckWait(waitTime),
		stan.SetManualAckMode(),
		stan.DurableName("orders-durable-subscriber"),
	); err != nil {
		h.logger.Fatal("Error when connecting to a channel", err)
	}

	return err
}
