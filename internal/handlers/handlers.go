package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/prashkotam/bednbreakfast/internal/config"
	"github.com/prashkotam/bednbreakfast/internal/forms"
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
	
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

//PostReservation handles posting of a reservation form

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName: r.Form.Get("last_name"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	//form.Has("first_name", r)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")


	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
	})
	return
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}	

//

func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {

	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("CAnnot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})


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
