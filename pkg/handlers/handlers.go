package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/esmaeilmirzaee/bookings/pkg/forms"
	"log"
	"net/http"

	"github.com/esmaeilmirzaee/bookings/pkg/config"
	"github.com/esmaeilmirzaee/bookings/pkg/models"
	"github.com/esmaeilmirzaee/bookings/pkg/renders"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (e *Repository) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (e *Repository) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{})
}

func (e *Repository) PostLoginPageHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Fprintf(w, "%+v %+v", email, password)
}


type data struct {
	OK bool `json="ok"`
	Message string `json="message"`
}

// JSONResponseHandler returns a valid json documents
func (e *Repository) JSONResponseHandler(w http.ResponseWriter, r *http.Request) {
	data := data{
		OK:true,
		Message: "This is a sample json.",
	}
	out, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		w.Write([]byte("Something went wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (e *Repository) SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "registration.page.tmpl", &models.TemplateData{})
}

func (e *Repository) RoomPageHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, r, "rooms.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostRoomPageHandler reserves a room for logged in user
func (e *Repository) PostRoomPageHandler(w http.ResponseWriter, r *http.Request) {
	err:=r.ParseForm()
	if err !=nil {
		log.Println(err)
		return
	}

	reservation:= models.ReservationForm{
		 FirstName: r.FormValue("first_name"),
		 LastName: r.FormValue("last_name"),
		 Phone: r.FormValue("phone"),
		 Email: r.FormValue("email"),
	}

	form := forms.New(r.PostForm)

	form.Has("first_name", r)

	if !form.Valid() {
		data:=make(map[string]interface{})
		data["reservation"] = reservation

		renders.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})

		return
	}
}

func (e *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["reservation"] = models.ReservationForm{}

	renders.RenderTemplate(w,r, "reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (e *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Cannot parse form")
		return
	}

	reservation := models.ReservationForm{
		FirstName: r.FormValue("first_name"),
		LastName: r.FormValue("last_name"),
		Email: r.FormValue("email"),
	}

	form := forms.New(r.PostForm)
	log.Println(form.Get("first_name"), reservation.FirstName, "->", r.MultipartForm)
	form.Required("first_name", "last_name", "email")
	form.CheckLength( 3, 50, r,"first_name", "last_name")
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		renders.RenderTemplate(w, r, "reservation.page.tmpl", &models.TemplateData{
			Data: data,
			Form: form,
		})

		return
	}

	log.Println("Everything is OK")
	e.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// ReservationSummary see details of reservation
func (e *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := e.App.Session.Get(r.Context(), "reservation").(models.ReservationForm)
	if !ok {
		log.Println("Cannot get reservation")
		e.App.Session.Put(r.Context(), "error", "Cannot get reservation")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data := make(map[string]interface{})
	data["reservation"] = reservation

	renders.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Data: data,
	})
}