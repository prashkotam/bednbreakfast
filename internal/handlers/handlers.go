package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/prashkotam/bednbreakfast/internal/config"
	"github.com/prashkotam/bednbreakfast/internal/models"
	"github.com/prashkotam/bednbreakfast/internal/render"
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

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		StringMap: stringmap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	var stringmap = map[string]string{}
	remoteip := m.App.Session.GetString(r.Context(), "remote_ip")

	stringmap["remote_ip"] = remoteip

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringmap,
	})
}

//Reservation renders the make a reservation page and displays form

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

//Generals renders the room page

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

//Majors renders the room page

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

//Availability renders the room page

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

//PostAvailability renders the room page

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	
	w.Write([]byte(fmt.Sprintf("Start date is %s and the end date is %s", start, end)))
}


type jsonResponse struct {
	OK 		bool 	`json: "ok"`
	Message string 	`json: "message"`

}

//AvailabilityJSON handles request for availability and sends json response

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {

	resp := jsonResponse {
		OK: false,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-type", "application/json" )
	w.Write(out)

}






//Contact renders the contact page

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}