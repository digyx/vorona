package handler

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/cbroglie/mustache"
	"github.com/digyx/vorona/internal"
	"github.com/gomarkdown/markdown"
)

// This transforms the internal.Book struct into a map used by mustache to fill in templates
func BookToMustache(book *internal.Book) map[string]interface{} {
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

// Combine base.html with pageName.html and its .style.html files
// context is the data used to fill in the template
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

	return mustache.RenderFile("templates/base.html", map[string]interface{}{
		"title":   title,         // <head><title>
		"content": content,       // pageName.html
		"style":   string(style), // pageName.style.html
	})
}
