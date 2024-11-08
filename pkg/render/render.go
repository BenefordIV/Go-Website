package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Beneford_IV/udemy_app/pkg/config"
	"github.com/Beneford_IV/udemy_app/pkg/models"
)

var funjctinMap = template.FuncMap{}
var app *config.AppConfig

// Constructor for new Template Package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("error template not in cache: ")
		return
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, td)

	//render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	//get all files named *.page.gohtml
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		log.Println("error parsing template: ", err)
		return cache, err
	}

	//loop through all pages
	for _, page := range pages {
		name := filepath.Base(page)
		log.Println("Page is currently ", page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println("error parsing template: ", err)
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			log.Println("error parsing template: ", err)
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				log.Println("error parsing template: ", err)
				return cache, err
			}
		}
		//add page to cache
		cache[name] = ts
	}
	return cache, nil
}
