package api

import (
	"context"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
	"wb/internal/api/order"
	"wb/pkg/logger"
)

func Module() fx.Option {
	return fx.Module(
		"Api Module",
		fx.Provide(
			fx.Annotate(setupMainRouter, fx.ParamTags(`group:"routes"`)),
			asRoute(order.New),
		),
		fx.Invoke(startServer),
	)
}

func startServer(lc fx.Lifecycle, router *chi.Mux, logger *logger.Logger) {
	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			if err := server.ListenAndServe(); err != nil {
				logger.Fatal("Cannot start server", err)
				return err
			}
			return nil
		},
		OnStop: func(context context.Context) error {
			logger.Info("Shutting down HTTP server")
			return server.Shutdown(context)
		},
	})
}

func setupMainRouter(routers []route) *chi.Mux {
	mainRouter := chi.NewRouter()
	for _, router := range routers {
		mainRouter.Mount("/", router.Routes())
	}
	return mainRouter
}
