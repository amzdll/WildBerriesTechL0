package logger

import (
	"wb/pkg/logger"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module(
		"logger",
		fx.Provide(logger.New),
	)
}
