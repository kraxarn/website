package db

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kraxarn/website/config"
)

var pool *pgxpool.Pool

func Connect() error {
	url, err := config.DatabaseUrl()
	if err != nil {
		return err
	}

	if pool != nil {
		return errors.New("pool already initialised")
	}

	pool, err = pgxpool.New(context.Background(), url)
	if err != nil {
		return err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	if pool == nil {
		return
	}

	pool.Close()
	pool = nil
}

func Acquire() (*pgxpool.Conn, error) {
	if pool == nil {
		return nil, errors.New("no connection pool available")
	}

	return pool.Acquire(context.Background())
}
