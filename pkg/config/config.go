package config

import "html/template"

// AppConfig - stores application configuration
type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
}
