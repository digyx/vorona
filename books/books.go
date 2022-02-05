package books

import "time"

type Book struct {
	ID          string
	Title       string
	ReleaseTime int64
}

func (book *Book) ToMap() map[string]interface{} {
	release_date := time.Unix(book.ReleaseTime, 0).UTC()
	is_released := time.Now().After(release_date)
	return map[string]interface{}{
		"book.id":           book.ID,
		"book.title":        book.Title,
		"book.release_date": release_date.Format("January 02 2006"),
		"book.is_released":  is_released,
	}
}
