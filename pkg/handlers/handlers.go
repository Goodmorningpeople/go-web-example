package handlers

import (
	"net/http"

	"github.com/Goodmorningpeople/learning_web_with_go/pkg/handlers/config"
	"github.com/Goodmorningpeople/learning_web_with_go/pkg/handlers/render"
)
// Holds data to be passed to a template, pass as parameter to a handler func 
	type TemplateData struct {
	StringMap map[string]string
	IntMap map[int]int
	Floatmap map[float32]float32
	Data map[string]interface{}
	CSRFToken string	
	Flash string
	Warning string
	Error string
}

// Repository pattern (very, very cool)
var Repo *Respository

type Respository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Respository {
	return &Respository{
		App: a,
	}
}

func NewHandler(r *Respository) {
	Repo = r
}

// Handler funcs that handle the rendered templates and write them to a server (response writer w is passed when used in HandleFunc() for this)
func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
