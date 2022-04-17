package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

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

func (self *context) apiBooks(w http.ResponseWriter, r *http.Request) {
	books, err := self.db.AllBooks()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}
