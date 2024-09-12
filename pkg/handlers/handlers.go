package handlers

import (
	"net/http"

	"github.com/Goodmorningpeople/go_web_example/pkg/config"
	"github.com/Goodmorningpeople/go_web_example/pkg/models"
	"github.com/Goodmorningpeople/go_web_example/pkg/render"
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
	remoteIp := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Respository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := map[string]string{
		"explanation": "Go is a really robust language for both systems programming (like rust but better) and web programming (arguably like javascript but better).", 
		"remote_ip": remoteIp,
	}

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}	