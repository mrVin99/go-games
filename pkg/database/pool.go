package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

var connPool *PG

func Conn() DB {
	return connPool
}

func init() {
	connPool = &PG{postgresPool()}
	migrate()
}

func postgresPool() *pgxpool.Pool {
	pgAddr := func() string {
		var (
			addr       = os.Getenv("PG_URI")
			defaultUrl = "postgres://postgres:postgres@localhost:5432/postgres"
		)

		if addr == "" {
			addr = defaultUrl
		}

		return addr
	}

	config, err := pgxpool.ParseConfig(pgAddr())
	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	if err = pool.Ping(context.Background()); err != nil {
		panic(err)
	}

	log.Println("Connected to postgres")

	return pool
}
