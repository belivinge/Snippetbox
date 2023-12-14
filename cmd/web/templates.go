package main

import (
	"html/template"
	"path/filepath"

	"github.com/belivinge/Snippetbox/pkg/models"
)

type templateData struct {
	Snippet *models.Snippet
	// snippet field in the templatedata struct
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// initialize a new map as a cache
	cache := map[string]*template.Template{}
	// filepath.Glob function to get a slice of all filepaths with page.html extension
	pages, err := filepath.Glob(filepath.Join(dir, "*page.html"))
	if err != nil {
		return nil, err
	}
	// loop through the pages one by one
	for _, page := range pages {
		// extract filename
		name := filepath.Base(page)
		// parse the page template file to template set
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}
		// parseglob method to add any layout templates
		ts, err = ts.ParseGlob(filepath.Join(dir, "*layout.html"))
		if err != nil {
			return nil, err
		}

		// parseglob method to add any partial templates
		ts, err = ts.ParseGlob(filepath.Join(dir, "*partial.html"))
		if err != nil {
			return nil, err
		}
		// add template set to cache using the name of the template
		cache[name] = ts
	}
	return cache, nil
}
