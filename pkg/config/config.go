package config

import "text/template"

// App config for project wide var to allow for easy config  and storage of project data eg. template caching 
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}
