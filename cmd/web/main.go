package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Goodmorningpeople/go_web_example/pkg/config"
	"github.com/Goodmorningpeople/go_web_example/pkg/handlers"
	"github.com/Goodmorningpeople/go_web_example/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app = config.AppConfig{}
var session *scs.SessionManager
func main()  {
	// Important!!! Set to true when in production for security
	app.InProduction = false

	// creating session and storing it in app config 
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction 
	app.Session = session

	// useCache set to false causes render function to create template cache every time render function is run but is slower, set to true for real users
	app.UseCache = app.InProduction
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("\nfatal error creating cache:", err)
	}

	// use global config in handlers.go and middleware.go
	handlers.NewHandler(handlers.NewRepo(&app))

	// use global config in render.go and assign template cache tc var to use in render.go
	app.TemplateCache = tc
	render.NewTemplates(&app)
	app.TemplateCache = tc

	fmt.Printf("Starting application on port %s", portNumber)

	// use http server with routes from routes.go
	svr := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	// serve http server 221
	err = svr.ListenAndServe()
	if err != nil {
		log.Fatal("\n fatal error serving http server")
	}
}
