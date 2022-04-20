package internal

// Generalized Boook struct that's used for passing data between modules
type Book struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseTime int64  `json:"release_time"`
}

// Generalized DB struct used for switching between sqlite and postgres (TBA)
type Database interface {
	GetBook(id string) (Book, error)
	GetAllBooks() ([]Book, error)
}
