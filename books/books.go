package books

import (
	"database/sql"
	"fmt"
	"time"
)

type Book struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseTime int64  `json:"release_time"`
}

func (book *Book) ToMustache() map[string]interface{} {
	release_date := time.Unix(book.ReleaseTime, 0).UTC()
	is_released := time.Now().After(release_date)
	return map[string]interface{}{
		"slug":         book.Slug,
		"title":        book.Title,
		"description":  book.Description,
		"release_date": release_date.Format("January 02 2006"),
		"is_released":  is_released,
	}
}

// Database
func GetAllBooks(db *sql.DB) ([]Book, error) {
	rows, err := db.Query(`
        SELECT slug, title, description, release_time
        FROM books
        ORDER BY release_time DESC`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	books := []Book{}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Slug, &book.Title, &book.Description, &book.ReleaseTime)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func GetBook(db *sql.DB, id string) (Book, error) {
	res := db.QueryRow(`
        SELECT slug, title, description, release_time
        FROM books
        WHERE slug=$1`, id)

	if res == nil {
		return Book{}, fmt.Errorf("error: could not find book id=%s", id)
	}

	var book Book
	res.Scan(&book.Slug, &book.Title, &book.Description, &book.ReleaseTime)

	return book, nil
}
