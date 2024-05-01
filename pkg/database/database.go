package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	DB interface {
		Exec(query string, args ...interface{}) error
		QueryRow(query string, args ...any) pgx.Row
		Query(query string, args ...any) (pgx.Rows, error)
	}

	PG struct {
		*pgxpool.Pool
	}
)

func (pg *PG) QueryRow(query string, args ...any) pgx.Row {
	return pg.Pool.QueryRow(context.Background(), query, args...)
}

func (pg *PG) Query(query string, args ...any) (pgx.Rows, error) {
	return pg.Pool.Query(context.Background(), query, args)
}

func (pg *PG) Exec(query string, args ...interface{}) error {
	_, err := pg.Pool.Exec(context.Background(), query, args...)
	return err
}
