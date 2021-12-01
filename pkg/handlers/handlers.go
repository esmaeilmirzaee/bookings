package handlers

import (
	"fmt"
	"net/http"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome</h1>")
}
