package main

import (
	"encoding/gob"
	"github.com/esmaeilmirzaee/bookings/internal/drivers"
	"github.com/esmaeilmirzaee/bookings/internal/handlers"
	"github.com/esmaeilmirzaee/bookings/internal/helpers"
	"github.com/esmaeilmirzaee/bookings/internal/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/esmaeilmirzaee/bookings/internal/config"
	"github.com/esmaeilmirzaee/bookings/internal/renderer"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

const dsn = "host=192.168.101.2 port=5234 dbname=bookings user=pgdmn password=secret"

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *drivers.DB) {
		err := db.SQL.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	log.Println("App is running on" + portNumber)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func run() (*drivers.DB, error) {
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

	// Connect to database
	log.Println("Connecting to database")
	db, err := drivers.ConnectSQL(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database; Bye")
	}
	log.Println("Connected to database.")

	app.UseCache = false
	tc, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc

	// Set loggers
	app.InfoLog = log.New(os.Stdout, "Info\t", log.Ldate|log.Ltime)
	app.ErrorLog  = log.New(os.Stderr, "Error\t\t", log.Ldate|log.Ltime|log.Lshortfile)

	repo := handlers.NewRepository(&app, db)
	handlers.NewHandlers(repo)
	renderer.NewTemplate(&app)
	helpers.NewHelpers(&app)

	return db, nil
}