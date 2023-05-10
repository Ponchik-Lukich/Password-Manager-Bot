package database

import (
	"context"
	"fmt"
	"log"
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

func LocalInit() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	folder := fmt.Sprintf("%s/%s", dir, "assets")
	tables := []string{"users", "services", "drop_constraint", "add_constraint"}

	for _, table := range tables {
		sql, err := os.ReadFile(fmt.Sprintf("%s/%s.sql", folder, table))
		if err != nil {
			log.Printf("Failed to run sql: %v", err)
			return err
		}

		_, err = pool.Exec(context.Background(), string(sql))
		if err != nil {
			log.Printf("Failed to create table: %v", err)
			return err
		}
	}
	return nil
}
