package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig - stores application configuration
type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	IsProduction  bool
	Session       *scs.SessionManager
}
