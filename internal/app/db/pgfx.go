package db

import (
	"context"
	"fmt"
	"log"
	"wb/internal/config"
	"wb/pkg/logger"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

func pgModule() fx.Option {
	return fx.Module(
		"PostgreSQL",
		fx.Provide(
			createPool,
			config.NewPgConfig,
		),
		fx.Invoke(closePool),
	)
}

func createPool(config *config.PgConfig, log *logger.Logger) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dsn := fmt.Sprintf(
		config.DsnTemplate,
		config.Host,
		config.Port,
		config.Name,
		config.User,
		config.Password,
		config.SslMode,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatal("Failed to connect to database.", err)
		return nil, err
	}

	log.Info("Database (PostgreSql) connection pool has been opened.")

	return pool, nil
}

func closePool(lc fx.Lifecycle, pool *pgxpool.Pool) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			pool.Close()
			log.Print("Database connection pool has been closed.")
			return nil
		},
	})
}
