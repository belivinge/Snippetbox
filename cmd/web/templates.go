package main

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/belivinge/Snippetbox/pkg/forms"
	"github.com/belivinge/Snippetbox/pkg/models"
)

// creates a nicley formatted string representation of a time.Time obkect
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// string-keyed map
var functions = template.FuncMap{
	"humanDate": humanDate,
}

type templateData struct {
	CurrentYear int
	// Error fields to the templateData struct
	// FormData   url.Values //the same underlying type as the r.PostForm map
	// FormErrors map[string]string
	Form    *forms.Form
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
		// template.FuncMap must be registered before the ParseFiles() method

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
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
