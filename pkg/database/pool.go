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
	log.Println("Connecting to database...")
	connPool = &PG{connect()}

	migrate()
}

func connect() *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(pgAddr())
	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")

	return pool
}

func pgAddr() string {
	var (
		addr       = os.Getenv("PG_URI")
		defaultUrl = "postgres://postgres:postgres@localhost:5432/postgres"
	)

	if addr == "" {
		addr = defaultUrl
	}

	return addr
}
