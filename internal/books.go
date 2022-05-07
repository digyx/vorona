package internal

// Generalized Boook struct that's used for passing data between modules
type Book struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseTime int64  `json:"release_time"`
	ImageURL    string `json:"image_url"`
}

// Generalized DB interface
type Database interface {
	GetBook(id string) (Book, error)
	GetAllBooks() ([]Book, error)
}
