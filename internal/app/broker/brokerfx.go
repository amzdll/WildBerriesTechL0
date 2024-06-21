package broker

import (
	nats_streaming "wb/internal/app/broker/nats-streaming"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"Message broker",
		nats_streaming.Module(),
	)
}
