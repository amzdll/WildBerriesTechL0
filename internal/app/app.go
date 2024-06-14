package app

import (
	"go.uber.org/fx"
	"wb/internal/app/api"
	"wb/internal/app/broker"
	"wb/internal/app/core"
	"wb/internal/app/db"
	"wb/internal/app/logger"
)

func Create() *fx.App {
	return fx.New(
		logger.Module(),

		db.Module(),
		core.Module(),

		api.Module(),
		broker.Module(),

		fx.NopLogger,
	)
}
