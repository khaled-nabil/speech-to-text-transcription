package postgres

import (
	"context"
	"fmt"
	"transcription-service/domain/persistance"
	"transcription-service/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func New(config *config.Config) (persistance.TranscriptionRepository, error) {
	ctx := context.Background()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Postgres.User,
		config.Postgres.Password,
		config.Postgres.Host,
		config.Postgres.Port,
		config.Postgres.Database,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Repository{
		pool: pool,
		ctx:  ctx,
	}, nil
}

func (r *Repository) Close() error {
	if r.pool != nil {
		r.pool.Close()
	}
	return nil
}
