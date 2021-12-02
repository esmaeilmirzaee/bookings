package handlers

import (
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
	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (e *Repository) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{})
}

func (e *Repository) SignupPageHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "registration.page.tmpl", &models.TemplateData{})
}
