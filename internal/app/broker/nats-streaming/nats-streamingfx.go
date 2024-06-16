package nats_streaming

import (
	"context"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"go.uber.org/fx"
	"wb/internal/broker/nats-streaming/order"
	"wb/pkg/logger"
)

func Module() fx.Option {
	return fx.Module(
		"Nats Streaming module",
		fx.Provide(
			asHandler(order.New),
		),
		fx.Invoke(fx.Annotate(setupConsumer, fx.ParamTags(`group:"topics"`))),
		fx.Provide(openConn),
		fx.Invoke(closeConn),
	)
}

func openConn(log *logger.Logger) (stan.Conn, error) {
	sc, err := stan.Connect(
		"test-cluster",
		"consumer",
		stan.NatsURL(nats.DefaultURL),
	)
	if err != nil {
		log.Fatal("Failed to connect to nats-streaming.", err)
		return nil, err
	}
	log.Info("Nats-Streaming connection has been opened.")
	return sc, err
}

func closeConn(lc fx.Lifecycle, connection stan.Conn, log *logger.Logger) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			if err := connection.Close(); err != nil {
				return err
			}
			log.Info("Nats Streaming connection closed.")
			return nil
		},
	})
}

func setupConsumer(handlers []mHandler, sc stan.Conn, log *logger.Logger) error {
	for _, handler := range handlers {
		if err := handler.Subscribe(sc); err != nil {
			log.Fatal("", err)
			return err
		}
	}
	return nil
}
