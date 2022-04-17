package db

import "github.com/digyx/vorona/internal"

type Database interface {
	Book(id string) (internal.Book, error)
	AllBooks() ([]internal.Book, error)
}
