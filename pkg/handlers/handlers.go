package handlers

import (
	"github.com/dar316/bookings/pkg/config"
	"github.com/dar316/bookings/pkg/models"
	"github.com/dar316/bookings/pkg/render"
	"net/http"
)

var Repo *Repository

// Repository Typ repozytorium
type Repository struct {
	App *config.AppConfig
}

// NewRepo Tworzy nowe repozytorium
func NewRepo(a *config.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

// NewHandlers Ustawia repozytorium dla handlers
func NewHandlers(r *Repository) {

	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// dodanie logiki

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// wysłanie danych do template

	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
