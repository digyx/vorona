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
