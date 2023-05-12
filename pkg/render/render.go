package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/AssassinAsh/going-go/pkg/cache"
	"github.com/AssassinAsh/going-go/pkg/config"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// RenderTemplate - fetched templates from the cache and renders
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	var ts map[string]*template.Template

	if app.UserCache {
		ts = app.TemplateCache
	} else {
		ts, _ = cache.CreateTemplateCache()
	}

	t, ok := ts[tmpl]

	if !ok {
		log.Fatal("Template unavailable")
	}

	buff := new(bytes.Buffer)

	err := t.Execute(buff, nil)

	if err != nil {
		log.Fatal(err)
	}

	_, err = buff.WriteTo(w)

	if err != nil {
		fmt.Println("Error occurred while executing the template: ", err)
	}
}
