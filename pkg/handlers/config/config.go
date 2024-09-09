package config

import "text/template"

// App config caontains the app config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}
