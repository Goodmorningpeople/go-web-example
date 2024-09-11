package main

import (
	"net/http"

	"github.com/Goodmorningpeople/learning_web_with_go/pkg/config"
	"github.com/Goodmorningpeople/learning_web_with_go/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// routes func contains all the routes, which are served to an http server in main.go (using chi router)
func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
