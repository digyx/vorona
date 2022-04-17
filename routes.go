package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cbroglie/mustache"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/digyx/vorona/internal"
)

func service() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	// Server Routes
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	r.Get("/", index)
	r.Get("/about", about)
	r.Get("/license", license)
	r.Get("/book/{slug:[A-Za-z]+}", book)

	// API Routes
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.SetHeader("content-type", "application/json"))

		r.Get("/book", apiBooks)
		r.Get("/book/{slug:[A-Za-z]+}", apiBook)
	})

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	books, err := internal.GetAllBooks(db)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var mustacheBooks []map[string]interface{}

	for _, elem := range books {
		mustacheBooks = append(mustacheBooks, elem.ToMustache())
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
		fmt.Println(err)
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
		fmt.Println(err)
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
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(res))
}

func book(w http.ResponseWriter, r *http.Request) {
	bookSlug := chi.URLParam(r, "slug")

	book, err := internal.GetBook(db, bookSlug)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Render the HTML Page
	res, err := renderTemplate(
		"book",
		book.Title,
		book.ToMustache(),
	)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(res))
}

// API Handlers
func apiBook(w http.ResponseWriter, r *http.Request) {
	bookSlug := chi.URLParam(r, "slug")
	book, err := internal.GetBook(db, bookSlug)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func apiBooks(w http.ResponseWriter, r *http.Request) {
	books, err := internal.GetAllBooks(db)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

//  Helper Functions
func renderTemplate(pageName string, title string, context interface{}) (string, error) {
	html_filename := fmt.Sprintf("templates/%s.html", pageName)
	style_filename := fmt.Sprintf("templates/%s.style.html", pageName)

	content, err := mustache.RenderFile(html_filename, context)
	if err != nil {
		return "", err
	}

	style, err := ioutil.ReadFile(style_filename)
	if err != nil {
		return "", err
	}

	res, err := mustache.RenderFile("templates/base.html", map[string]interface{}{
		"title":   title,
		"content": content,
		"style":   string(style),
	})
	if err != nil {
		return "", err
	}

	return res, nil
}
