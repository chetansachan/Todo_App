package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Printf("Unable to parse DATABASE_URL: %v", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("Unable to connect to database: %v", err)
		return nil, err
	}

	err2 := pool.Ping(ctx)
	if err2 != nil {
		log.Printf("Unable to ping database: %v", err)
		return nil, err
	}

	log.Printf("Database connected to: %v", databaseURL)
	return pool, nil
}
