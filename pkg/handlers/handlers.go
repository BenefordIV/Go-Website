package handlers

import (
	"net/http"

	"github.com/Beneford_IV/udemy_app/pkg/config"
	"github.com/Beneford_IV/udemy_app/pkg/models"
	"github.com/Beneford_IV/udemy_app/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// All Handler functions require TWO arguements. A ResponseWriter and a Request
func (h *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About is the about page handler
func (h *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	td := models.TemplateData{
		StringMap: stringMap,
	}

	// send the data to the template
	render.RenderTemplate(w, "about.page.gohtml", &td)
}
