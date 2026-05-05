package postgres

import (
	"context"
	"fmt"
	"mikhaylovilya/map-sorcerer/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDBPool(ctx context.Context, config *config.Config) (*pgxpool.Pool, error) {
	const op = "infra.postgres.NewPool"
	dbpool, err := pgxpool.New(ctx, config.PostgresConfig.URI)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return dbpool, nil
}
