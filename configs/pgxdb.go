package configs

import (
	"context"

	"github.com/engine/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPg(c models.PostgresqlConfig) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), c.URL)
	if err != nil {
		return nil, err
	}
	return dbPool, nil
}
