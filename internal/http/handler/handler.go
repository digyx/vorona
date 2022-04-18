package handler

import (
	"net/http"

	"github.com/digyx/vorona/internal"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type context struct {
	db internal.Database
}

func New(db internal.Database) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	ctx := context{db: db}

	// Server Routes
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
	})

	r.Get("/", ctx.index)
	r.Get("/about", about)
	r.Get("/license", license)
	r.Get("/book/{slug:[A-Za-z]+}", ctx.book)

	// API Routes
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.SetHeader("content-type", "application/json"))

		r.Get("/book", ctx.apiBooks)
		r.Get("/book/{slug:[A-Za-z]+}", ctx.apiBook)
	})

	return r
}
