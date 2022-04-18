package sqlite

import (
	"fmt"

	"github.com/digyx/vorona/internal"
)

// This is the actual struct stored in SQLite
// Currently unused...not great
type bookSchema struct {
	ID          int64  `db:"book_id"`
	Slug        string `db:"slug"`
	Title       string `db:"title"`
	Description string `db:"description"`
	ReleaseTime int64  `db:"release_time"`
}

// Grab all the books in a list; ordered from newest to oldest release time
func (self *SQLite) AllBooks() ([]internal.Book, error) {
	rows, err := self.DB.Query(`
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

// Grab just a single book by its slug
func (self *SQLite) Book(slug string) (internal.Book, error) {
	res := self.DB.QueryRow(`
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
