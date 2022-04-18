package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

// Implements internal.Database
type SQLite struct {
	DB *sql.DB
}

// Connect to the SQLite database
func New(path string) (*SQLite, error) {
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return &SQLite{DB: conn}, nil
}
