package api

import (
	"context"
	"net/http"
	"wb/internal/api/order"
	"wb/internal/api/swagger"
	"wb/internal/config"
	"wb/pkg/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.uber.org/fx"

	_ "wb/api"
)

func Module() fx.Option {
	return fx.Module(
		"Api Module",
		fx.Provide(
			fx.Annotate(setupMainRouter, fx.ParamTags(`group:"routes"`)),
			asRoute(order.New),
			config.NewApiConfig,
		),
		fx.Invoke(startServer),
	)
}

func startServer(
	lc fx.Lifecycle,
	router *chi.Mux,
	config *config.ApiConfig,
	logger *logger.Logger,
) {
	server := &http.Server{
		Addr:    config.Host + ":" + config.Port,
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

func setupMainRouter(routers []route, stage string) *chi.Mux {
	mainRouter := chi.NewRouter()
	cacheTime := 300
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           cacheTime,
	})
	mainRouter.Use(corsMiddleware.Handler)

	if stage != "prod" {
		mainRouter.Mount(swagger.Routes())
	}

	for _, router := range routers {
		mainRouter.Mount(router.Routes())
	}

	return mainRouter
}
