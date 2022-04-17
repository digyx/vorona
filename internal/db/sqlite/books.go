package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/digyx/vorona/internal"
)

type bookSchema struct {
	ID          int64  `db:"book_id"`
	Slug        string `db:"slug"`
	Title       string `db:"title"`
	Description string `db:"description"`
	ReleaseTime int64  `db:"release_time"`
}

type SQLite struct {
	DB *sql.DB
}

func (self *SQLite) AllBooks() ([]internal.Book, error) {
	rows, err := self.DB.Query(`
		SELECT slug, title, description, release_time
		FROM books
		ORDER BY release_time DESC`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

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

func (self *SQLite) Book(id string) (internal.Book, error) {
	res := self.DB.QueryRow(`
		SELECT slug, title, description, release_time
		FROM books
		WHERE slug=$1`, id)

	if res == nil {
		return internal.Book{}, fmt.Errorf("error: could not find book id=%s", id)
	}

	var book internal.Book
	res.Scan(&book.Slug, &book.Title, &book.Description, &book.ReleaseTime)

	return book, nil
}
