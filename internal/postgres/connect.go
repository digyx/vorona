package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Postgres struct {
	DB *pgxpool.Pool
}

func New(path string) (*Postgres, error) {
	conn, err := pgxpool.Connect(context.Background(), path)
	if err != nil {
		return nil, err
	}

	return &Postgres{DB: conn}, nil
}
