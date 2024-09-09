package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Goodmorningpeople/learning_web_with_go/pkg/handlers/config"
)
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a 	
}
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc := map[string]*template.Template{}
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()

	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("\nfatal error accessing cache")
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)
	if err != nil {
		fmt.Print("\nerror executing buffer:", err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Print("\nerror writing into buffer:", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		fmt.Print("\nerror accessing files:", err)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}	

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl") 
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts 
	}
	return myCache, nil
}