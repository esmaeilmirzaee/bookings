package handlers

import (
	"encoding/json"
	"fmt"
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
