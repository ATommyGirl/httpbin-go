package main

import (
	"net/http"
	"strconv"
)

func procRedirectTo(w http.ResponseWriter, r *http.Request) {

	statusCode, err := strconv.Atoi(r.FormValue("status_code"))

	if err != nil {
		statusCode = 302
	}

	redirectTo(w, statusCode, r.FormValue("url"))
}

func redirectTo(responseWriter http.ResponseWriter, statusCode int, redirectToLocation string) {
	if statusCode < 300 || statusCode > 399 {
		statusCode = 302
	}
	responseWriter.Header().Set("Location", redirectToLocation)
	responseWriter.WriteHeader(statusCode)
}
