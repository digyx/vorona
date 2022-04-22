package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "modernc.org/sqlite"
)

// Implements internal.Database
type SQL struct {
	DB *sql.DB
}

// Connect to the SQLite database
func New(driver, path string) (*SQL, error) {
	conn, err := sql.Open(driver, path)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Ensure that the connection doesn't error out

	if err := conn.PingContext(ctx); err != nil {
		return nil, err
	}

	return &SQL{DB: conn}, nil
}
