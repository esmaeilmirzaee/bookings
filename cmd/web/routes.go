package main

import (
	"net/http"

	"github.com/esmaeilmirzaee/bookings/pkg/config"
	"github.com/esmaeilmirzaee/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()

	mux.Get("/", handlers.HomePageHandler)

	return mux
}
