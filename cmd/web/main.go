package main

import (
	"log"
	"net/http"

	"github.com/esmaeilmirzaee/bookings/pkg/config"
	"github.com/esmaeilmirzaee/bookings/pkg/renders"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {
	app.IsProduction = false

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
