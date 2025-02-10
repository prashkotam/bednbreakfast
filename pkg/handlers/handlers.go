package handlers

import (
	"net/http"

	"github.com/prashkotam/bednbreakfast/pkg/config"
	"github.com/prashkotam/bednbreakfast/pkg/models"
	"github.com/prashkotam/bednbreakfast/pkg/render"
)

type Repository struct {
	App *config.Appconfig
}

var Repo *Repository

func NewRepo(a *config.Appconfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

// handler functions

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	var stringmap = map[string]string{}
	stringmap["test1"] = "Hello world!!!"

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{
		StringMap: stringmap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	var stringmap = map[string]string{}
	remoteip := m.App.Session.GetString(r.Context(), "remote_ip")

	stringmap["remote_ip"] = remoteip

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringmap,
	})
}
