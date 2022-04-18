package internal

type Book struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseTime int64  `json:"release_time"`
}

type Database interface {
	Book(id string) (Book, error)
	AllBooks() ([]Book, error)
}
