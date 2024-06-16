package api

import (
	"context"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/fx"
	"net/http"
	"wb/internal/api/order"
	"wb/pkg/logger"

	_ "wb/api"
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

// @title          Chat API
// @version        1.0
// @description    Shop Api
//
// @host        127.0.0.1:8000
// @BasePath    /
func setupMainRouter(routers []route) *chi.Mux {
	mainRouter := chi.NewRouter()
	mainRouter.Get(
		"/docs/*",
		httpSwagger.Handler(httpSwagger.URL("doc.json")),
	)
	for _, router := range routers {
		mainRouter.Mount("/", router.Routes())
	}
	return mainRouter
}
