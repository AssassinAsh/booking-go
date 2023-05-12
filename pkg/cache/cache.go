package cache

import (
	"fmt"
	"html/template"
	"path/filepath"
)

// CreateTemplateCache - Loads all the templates in the cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	fmt.Println("Templates Cached Successfully")

	return cache, nil
}
