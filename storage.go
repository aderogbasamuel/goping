package main

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage interface {
	Save(code string, url string) error
	Get(code string) (string, error)
}

type PostgresStorage struct {
	db *pgxpool.Pool
}

func (p PostgresStorage) Save(code string, url string) error {
	_, err := p.db.Exec(
		context.Background(),
		`INSERT INTO urls (code, original_url)
		 VALUES ($1, $2)`,
		code,
		url,
	)

	return err
}

func (p PostgresStorage) Get(code string) (string, error) {
	var url string

	err := p.db.QueryRow(
		context.Background(),
		`SELECT original_url
		 FROM urls
		 WHERE code = $1`,
		code,
	).Scan(&url)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errors.New("URL not found")
		}

		return "", err
	}

	return url, nil
}