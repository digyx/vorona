package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	DB *sql.DB
}

func New() (*SQLite, error) {
	conn, err := sql.Open("sqlite3", "./vorona.db")
	if err != nil {
		return nil, err
	}

	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}

	createStmt := `
	CREATE TABLE IF NOT EXISTS books (
		slug         TEXT    NOT NULL UNIQUE,
		title        TEXT    NOT NULL,
		description  TEXT    NOT NULL,
		release_time INTEGER NOT NULL
	)
	`
	_, err = tx.Exec(createStmt)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &SQLite{DB: conn}, nil
}
