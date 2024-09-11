package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// No surf is used for csrf protection 
func NoSurf(next http.Handler) http.Handler {
	crsfHandler := nosurf.New(next)

	crsfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/", 
		Secure: false, // set Secure to true in production
		SameSite: http.SameSiteLaxMode, 
	})

	return crsfHandler
}