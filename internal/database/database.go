package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func Connect() error {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil
	}

	var err error
	pool, err = pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		return err
	}

	return nil
}
