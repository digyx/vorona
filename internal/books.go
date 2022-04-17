package internal

import (
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
)

type Book struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseTime int64  `json:"release_time"`
}

func (book *Book) ToMustache() map[string]interface{} {
	release_date := time.Unix(book.ReleaseTime, 0).UTC()
	is_released := time.Now().After(release_date)

	description := string(markdown.ToHTML([]byte(book.Description), nil, nil))

	return map[string]interface{}{
		"slug":         book.Slug,
		"title":        book.Title,
		"description":  strings.TrimSpace(description),
		"release_date": release_date.Format("January 02 2006"),
		"is_released":  is_released,
	}
}
