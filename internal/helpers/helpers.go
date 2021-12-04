package helpers

import (
	"fmt"
	"github.com/esmaeilmirzaee/bookings/internal/config"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

// NewHelpers Sets app
func NewHelpers(a *config.AppConfig) {
	app = a
}

// ClientError logs client-related errors
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, "Client error with status of", status)
}

// ServerError responses with
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("Server error %+v %+v", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}