package handlers

import (
	"net/http"

	"github.com/AssassinAsh/going-go/pkg/config"
	"github.com/AssassinAsh/going-go/pkg/models"
	"github.com/AssassinAsh/going-go/pkg/render"
)

var Repo *Repository

type Repository struct {
	AppConfig *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		AppConfig: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	strings := map[string]string{
		"text": "Home page para",
	}

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: strings,
		Title:     "Home Page",
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	strings := map[string]string{
		"text": "About page para",
	}

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: strings,
		Title:     "About Page",
	})
}
