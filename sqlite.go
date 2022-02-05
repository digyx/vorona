package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func connectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./vorona.db")
	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}

	createStmt := `
	CREATE TABLE IF NOT EXISTS books (
		slug         TEXT    NOT NULL UNIQUE,
		title        TEXT    NOT NULL,
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

	return db, nil
}
