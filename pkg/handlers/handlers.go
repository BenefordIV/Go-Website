package handlers

import (
	"net/http"

	"github.com/Beneford_IV/udemy_app/pkg/render"
)

// All Handler functions require TWO arguements. A ResponseWriter and a Request
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.gohtml")
}
