package app

import (
	"flag"
	"wb/internal/app/api"
	"wb/internal/app/broker"
	"wb/internal/app/core"
	"wb/internal/app/db"
	"wb/internal/app/logger"
	"wb/internal/app/validator"
	"wb/internal/config"

	"go.uber.org/fx"
)

func Create() *fx.App {
	return fx.New(
		validator.Module(),
		logger.Module(),
		db.Module(),
		core.Module(),
		api.Module(),
		broker.Module(),
		fx.NopLogger,

		fx.Provide(
			parseStage,
			config.New,
		),
	)
}

func parseStage() string {
	var stage string

	flag.StringVar(&stage, "stage", "local", "Stage (local, dev, prod)")
	flag.Parse()

	return stage
}
