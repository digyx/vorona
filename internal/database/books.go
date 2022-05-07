package database

import (
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
	ImageURL    string `db:"image_url"`
}

// Grab all the books in a list; ordered from newest to oldest release time
func (self *SQL) GetAllBooks() ([]internal.Book, error) {
	rows, err := self.DB.Query(`
		SELECT slug, title, description, release_time, image_url
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
		err := rows.Scan(
			&result.Slug,
			&result.Title,
			&result.Description,
			&result.ReleaseTime,
			&result.ImageURL,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, result)
	}

	return books, nil
}

// Grab just a single book by its slug
func (self *SQL) GetBook(slug string) (internal.Book, error) {
	var book internal.Book
	err := self.DB.QueryRow(`
		SELECT slug, title, description, release_time, image_url
		FROM books
		WHERE slug=$1
	`, slug).Scan(
		&book.Slug,
		&book.Title,
		&book.Description,
		&book.ReleaseTime,
		&book.ImageURL,
	)

	if err != nil {
		return internal.Book{}, err
	}

	return book, nil
}
