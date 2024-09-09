package handlers

import (
	"net/http"

	"github.com/Goodmorningpeople/learning_web_with_go/pkg/handlers/config"
	"github.com/Goodmorningpeople/learning_web_with_go/pkg/handlers/render"
)

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
func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
