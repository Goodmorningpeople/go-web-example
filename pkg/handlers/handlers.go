package handlers

import (
	"net/http"

	"github.com/Goodmorningpeople/learning_web_with_go/pkg/config"
	"github.com/Goodmorningpeople/learning_web_with_go/pkg/models"
	"github.com/Goodmorningpeople/learning_web_with_go/pkg/render"
)

// repository pattern (very, very cool)
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

// handler funcs that handle the rendered templates and write them to a server with oppotional logic passed to template (response writer w is passed when used in HandleFunc() for this)
func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "Go is a really robust language for both systems programming (like rust but better) and web programming (arguably like javascript but better)."

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}	