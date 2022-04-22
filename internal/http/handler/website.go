package handler

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (self *context) index(w http.ResponseWriter, r *http.Request) {
	books, err := self.db.GetAllBooks()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var mustacheBooks []map[string]interface{}

	for _, book := range books {
		mustacheBooks = append(mustacheBooks, BookToMustache(&book))
	}

	// Render the HTML Page
	res, err := renderTemplate(
		"index",
		"Vorona",
		map[string]interface{}{
			"books": mustacheBooks,
		},
	)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(res))
}

func about(w http.ResponseWriter, r *http.Request) {
	res, err := renderTemplate(
		"about",
		"About Me - Vorona",
		make(map[string]interface{}),
	)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(res))
}

func license(w http.ResponseWriter, r *http.Request) {
	res, err := renderTemplate(
		"license",
		"License - Vorona",
		make(map[string]interface{}),
	)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(res))
}

func (self *context) book(w http.ResponseWriter, r *http.Request) {
	bookSlug := chi.URLParam(r, "slug")

	book, err := self.db.GetBook(bookSlug)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Render the HTML Page
	res, err := renderTemplate(
		"book",
		book.Title,
		BookToMustache(&book),
	)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(res))
}
