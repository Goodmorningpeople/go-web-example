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

// Creates a new template using create template func and executes them using a buffer 
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc := map[string]*template.Template{}
	// UseCache is defined in main so explanation is located there
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()

	}

	// Accessing template from template cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("\nfatal error accessing cache")
	}

	// Executes the template by using a buffer that contains bytes as the parameter for execute
	buf := new(bytes.Buffer)
	err := t.Execute(buf, nil)
	if err != nil {
		fmt.Print("\nerror executing buffer:", err)
	}

	// Too stupid to understand what is happening with the buffer, fix later
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Print("\nerror writing into buffer:", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Gets the files from the templates directory that have the extension page.tmpl (the main page files)
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		fmt.Print("\nerror accessing files:", err)
		return myCache, err
	}

	// ranges through all the templates collected from the pages var (page.tmpl files), parsing all the templates and adding to the cache
	for _, page := range pages {

		// .Base removes the file extensions so that the name is concise for more readable and writable code when calling this func 
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}	
		
		// need layout template so that html is properly represented in webpage, layout is accessed and stored in matches var 
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// if there is at least one layout, will parse the layout(s) so that html using the layout(s) is properly represented
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