package broker

import (
	"go.uber.org/fx"
	"wb/internal/app/broker/nats-streaming"
)

func Module() fx.Option {
	return fx.Module(
		"Message broker",
		nats_streaming.Module(),
	)
}
