package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// no surf is used for csrf protection
func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)

	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return crsfHandler
}

// loads and saves the session
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
