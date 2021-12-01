package handlers

import (
	"net/http"

	"github.com/esmaeilmirzaee/bookings/pkg/models"
	"github.com/esmaeilmirzaee/bookings/pkg/renders"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
