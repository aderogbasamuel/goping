package main

import (
    "context"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
)

func connectDB() (*pgxpool.Pool, error) {
    databaseURL := os.Getenv("DATABASE_URL")

    return pgxpool.New(context.Background(), databaseURL)
}