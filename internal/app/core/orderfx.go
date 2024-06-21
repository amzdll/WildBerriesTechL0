package core

import (
	api "wb/internal/api/order"
	broker "wb/internal/broker/nats-streaming/order"
	pgRepository "wb/internal/db/pg/order"
	redisRepository "wb/internal/db/redis/order"
	service "wb/internal/service/order"

	"go.uber.org/fx"
)

func orderModule() fx.Option {
	return fx.Module(
		"Order core module",
		fx.Provide(
			fx.Annotate(
				redisRepository.New,
				fx.As(new(service.RedisRepository)),
			),
			fx.Annotate(
				pgRepository.New,
				fx.As(new(service.PgRepository)),
			),
			fx.Annotate(
				service.New,
				fx.As(new(broker.Service)),
				fx.As(new(api.Service)),
			),
		),
	)
}
