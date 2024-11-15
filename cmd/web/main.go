package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Beneford_IV/udemy_app/pkg/config"
	"github.com/Beneford_IV/udemy_app/pkg/handlers"
	ren "github.com/Beneford_IV/udemy_app/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change to true when in production
	app.InProduction = false

	session = scs.New()
	//This will make the session expire after 24 hours
	session.Lifetime = 24 * time.Hour
	//This will make the session persist even after the browser is closed
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := ren.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	app.UseCache = true

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
