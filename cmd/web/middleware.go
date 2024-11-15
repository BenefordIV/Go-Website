package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF to all Post Requests
func NoSurf(next http.Handler) http.Handler {
	csrfhandler := nosurf.New(next)

	csrfhandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfhandler
}

// SessionLoad loads and saves the session on every request to a cookie
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
