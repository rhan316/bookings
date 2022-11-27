package render

import (
	"bytes"
	"github.com/dar316/bookings/pkg/config"
	"github.com/dar316/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplates Tworzy konfigurację dla pakietu template
func NewTemplates(a *config.AppConfig) {

	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// Template renders a template
func Template(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template form template cache")
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buff, td)

	_, err := buff.WriteTo(w)
	if err != nil {
		log.Println("error writing template to browser")
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	// Weź wszystkie pliki nazwane: *.page.tmpl z ./templates
	pages, err := filepath.Glob("./templates/*page.tmpl")
	if err != nil {
		return cache, err
	}

	// Przeszukaj wszystkie pliki z końcówką *.page.tmpl
	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {

			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")

			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}

	return cache, nil
}
