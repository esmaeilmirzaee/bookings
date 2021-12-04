package main

import (
	"encoding/gob"
	"github.com/esmaeilmirzaee/bookings/internal/handlers"
	"github.com/esmaeilmirzaee/bookings/internal/helpers"
	"github.com/esmaeilmirzaee/bookings/internal/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/esmaeilmirzaee/bookings/internal/config"
	"github.com/esmaeilmirzaee/bookings/internal/renders"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	log.Println("App is running on" + portNumber)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// Data to store in session
	gob.Register(models.Reservation{})

	// Alter the following to true in production model
	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Persist = false
	session.Cookie.Secure = app.IsProduction

	app.Session = session

	app.UseCache = false
	tc, err := renders.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return err
	}

	app.TemplateCache = tc

	// Set loggers
	app.InfoLog = log.New(os.Stdout, "Info\t", log.Ldate|log.Ltime)
	app.ErrorLog  = log.New(os.Stderr, "Error\t\t", log.Ldate|log.Ltime|log.Lshortfile)

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)
	renders.NewTemplate(&app)
	helpers.NewHelpers(&app)

	return nil
}