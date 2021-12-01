package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/esmaeilmirzaee/bookings/pkg/config"
	"github.com/esmaeilmirzaee/bookings/pkg/renders"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
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
	}

	app.TemplateCache = tc
	renders.NewTemplate(&app)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	log.Println("App is running on" + portNumber)
	srv.ListenAndServe()
}
