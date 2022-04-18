package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Return the raw JSON for a book given a slug
func (self *context) apiBook(w http.ResponseWriter, r *http.Request) {
	bookSlug := chi.URLParam(r, "slug")

	book, err := self.db.Book(bookSlug)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// Return a list of all books
func (self *context) apiBooks(w http.ResponseWriter, r *http.Request) {
	books, err := self.db.AllBooks()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}
