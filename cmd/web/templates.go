package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"snippetbox.rakesh.net/internal/models"
	"snippetbox.rakesh.net/ui"
	"time"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	// Get all page templates
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	// Loop over each page template
	for _, page := range pages {
		// Extract filename from the full path
		name := filepath.Base(page)

		//create a slice containing the filepath patterns for the templates we want to prove
		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		// Store the parsed template in the cache
		cache[name] = ts
	}

	// Return the cache containing parsed templates
	return cache, nil
}
