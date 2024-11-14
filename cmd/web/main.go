package main

import (
	"log"
	"net/http"

	"github.com/Beneford_IV/udemy_app/pkg/config"
	"github.com/Beneford_IV/udemy_app/pkg/handlers"
	ren "github.com/Beneford_IV/udemy_app/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := ren.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	ren.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	//_ = http.ListenAndServe(portNumber, nil)
}
