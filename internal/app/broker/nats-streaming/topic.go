package nats_streaming

import (
	"github.com/nats-io/stan.go"
	"go.uber.org/fx"
)

type mHandler interface {
	Subscribe(conn stan.Conn) error
}

func asHandler(f interface{}) interface{} {
	return fx.Annotate(
		f,
		fx.As(new(mHandler)),
		fx.ResultTags(`group:"topics"`),
	)
}
