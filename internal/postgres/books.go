package postgres

import (
	"context"
	"fmt"

	"github.com/digyx/vorona/internal"
)

type bookSchema struct {
	ID          int64
	Slug        string
	Title       string
	Description string
	ReleaseTime string
}

func (self *Postgres) GetAllBooks() ([]internal.Book, error) {
	rows, err := self.DB.Query(
		context.Background(), `
		SELECT slug, title, description, release_time
		FROM books
		ORDER BY release_time DESC
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Decode the rows into a list of internal.Book
	books := []internal.Book{}
	for rows.Next() {
		var result internal.Book
		err := rows.Scan(&result.Slug, &result.Title, &result.Description, &result.ReleaseTime)
		if err != nil {
			return nil, err
		}

		books = append(books, result)
	}

	return books, nil
}

func (self *Postgres) GetBook(slug string) (internal.Book, error) {
	res := self.DB.QueryRow(
		context.Background(), `
		SELECT slug, title, description, release_time
		FROM books
		WHERE slug=$1
	`, slug)

	if res == nil {
		return internal.Book{}, fmt.Errorf("error: could not find book with slug=%s", slug)
	}

	// Decode into internal.Book
	var book internal.Book
	res.Scan(&book.Slug, &book.Title, &book.Description, &book.ReleaseTime)

	return book, nil
}
