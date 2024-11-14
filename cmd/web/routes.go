package main

import (
	"net/http"

	"github.com/Beneford_IV/udemy_app/pkg/config"
	"github.com/Beneford_IV/udemy_app/pkg/handlers"
	"github.com/bmizerany/pat"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
