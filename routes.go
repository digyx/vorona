package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cbroglie/mustache"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"dev.vorona.gg/digyx/vorona/books"
)

func service() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	// Server Routes
	r.Get("/", index)

	return r
}

func index(w http.ResponseWriter, r *http.Request) {
	books, err := books.GetAllBooks(db)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var mustacheBooks []map[string]interface{}

	for _, elem := range books {
		mustacheBooks = append(mustacheBooks, elem.ToMustache())
	}

	// Render the HTML Template
	html, err := mustache.RenderFile("templates/index.html", map[string]interface{}{
		"books": mustacheBooks,
	})
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(html)

	// Load the CSS
	style, err := ioutil.ReadFile("templates/index.style.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Combine with base
	res, err := mustache.RenderFile("templates/base.html", map[string]interface{}{
		"title":   "Vorona",
		"content": html,
		"style":   string(style),
	})
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(res))
}
