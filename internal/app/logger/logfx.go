package logger

import (
	"go.uber.org/fx"
	"wb/pkg/logger"
)

func Module() fx.Option {
	return fx.Module(
		"logger",
		fx.Provide(logger.New),
	)
}
