package handler

import (
	"fmt"
	"io/ioutil"

	"github.com/cbroglie/mustache"
)

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
