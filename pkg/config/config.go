package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// app config for project wide var to allow for easy config  and storage of project data eg. template caching
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
	Session       *scs.SessionManager
}
