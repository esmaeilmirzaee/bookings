package main

import (
	"net/http"

	"github.com/esmaeilmirzaee/bookings/internal/config"
	"github.com/esmaeilmirzaee/bookings/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewMux()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.HomePageHandler)
	mux.Get("/login", handlers.Repo.LoginPageHandler)
	mux.Post("/login", handlers.Repo.PostLoginPageHandler)
	mux.Get("/json", handlers.Repo.JSONResponseHandler)
	mux.Get("/rooms", handlers.Repo.RoomPageHandler)
	mux.Post("/rooms", handlers.Repo.PostRoomPageHandler)
	mux.Get("/signup", handlers.Repo.SignupPageHandler)

	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Post("/reservation", handlers.Repo.PostReservation)

	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))
	return mux
}
