package core

import (
	"go.uber.org/fx"
	api "wb/internal/api/order"
	broker "wb/internal/broker/nats-streaming/order"
	repository "wb/internal/repository/order"
	service "wb/internal/service/order"
)

func orderModule() fx.Option {
	return fx.Module(
		"Order core module",
		fx.Provide(
			fx.Annotate(
				repository.New,
				fx.As(new(service.Repository)),
			),
			fx.Annotate(
				service.New,
				fx.As(new(broker.Service)),
				fx.As(new(api.Service)),
			),
		),
	)
}
