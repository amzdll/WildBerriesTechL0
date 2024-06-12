package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"log"
	"wb/pkg/logger"
)

func pgModule() fx.Option {
	return fx.Module(
		"PostgreSQL",
		fx.Provide(newPool),
		fx.Invoke(closePool),
	)
}

func newPool(log *logger.Logger) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		"localhost", "5432", "wb",
		"postgres", "postgres", "disable",
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
