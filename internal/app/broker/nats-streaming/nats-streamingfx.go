package nats_streaming

import (
	"context"
	"github.com/nats-io/stan.go"
	"go.uber.org/fx"
	"wb/internal/broker/nats-streaming/order"
	"wb/internal/config"
	"wb/pkg/logger"
)

func Module() fx.Option {
	return fx.Module(
		"Nats Streaming module",
		fx.Provide(
			openConn,
			asHandler(order.New),
			config.NewStanConfig,
		),
		fx.Invoke(
			fx.Annotate(setupConsumer, fx.ParamTags(`group:"topics"`)),
			closeConn,
		),
	)
}

func openConn(config *config.StanConfig, log *logger.Logger) (stan.Conn, error) {
	sc, err := stan.Connect(
		config.ClusterId,
		config.ClientId,
		stan.NatsURL(config.Host+":"+config.Port),
	)
	if err != nil {
		log.Fatal("Failed to connect to nats-streaming.", err)
		return nil, err
	}
	log.Info("Message broker (Nats-Streaming) connection has been opened.")
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
