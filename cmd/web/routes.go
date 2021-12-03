package main

import (
	"net/http"

	"github.com/esmaeilmirzaee/bookings/pkg/config"
	"github.com/esmaeilmirzaee/bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()

	mux.Get("/", handlers.Repo.HomePageHandler)
	mux.Get("/login", handlers.Repo.LoginPageHandler)
	mux.Post("/login", handlers.Repo.PostLoginPageHandler)
	mux.Get("/json", handlers.Repo.JSONResponseHandler)
	mux.Get("/rooms", handlers.Repo.RoomPageHandler)
	mux.Get("/signup", handlers.Repo.SignupPageHandler)

	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))
	return mux
}
