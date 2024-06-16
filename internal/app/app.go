package app

import (
	"context"
	"go.uber.org/fx"
	"log"
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

		fx.Invoke(manageLifecycle),
	)
}

func manageLifecycle(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: startApp,
		OnStop:  stopApp,
	})
}

func startApp(_ context.Context) error {
	log.Println("Starting application")
	return nil
}

func stopApp(_ context.Context) error {
	log.Println("Starting application")
	return nil
}
