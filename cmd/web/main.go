package main

import (
	"log"
	"net/http"

	"github.com/esmaeilmirzaee/bookings/pkg/config"
)

const portNumber = ":8080"

var app config.AppConfig

func main() {

	log.Println("App is running on" + portNumber)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	srv.ListenAndServe()
}
