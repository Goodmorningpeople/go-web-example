package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Goodmorningpeople/learning_web_with_go/pkg/config"
	"github.com/Goodmorningpeople/learning_web_with_go/pkg/handlers"
	"github.com/Goodmorningpeople/learning_web_with_go/pkg/render"
)

const portNumber = ":8080"

func main() {
	app := config.AppConfig{}

	// UseCache set to true causes render function to create template cache every time render function is run but is slower, set to false for real users
	app.UseCache = false
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("\nfatal error creating cache:", err)
	}

	// Use global config in handlers.go
	handlers.NewHandler(handlers.NewRepo(&app))

	// Use global config in render.go and assign template cache tc var to use in render.go
	app.TemplateCache = tc
	render.NewTemplates(&app)
	app.TemplateCache = tc

	fmt.Printf("Starting application on port %s", portNumber)

	// Use http server with routes from routes.go 
	svr := http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	// Serve http server 
	err = svr.ListenAndServe()
	if err != nil {
		log.Fatal("\n fatal error serving http server")
	}
}
