package controllers

import (
	"fmt"
	"net/http"
)

// IndexHandler serves our index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello from %v", r.RequestURI)
	if err != nil {
		fmt.Printf("Unable to complete request: %v\n", err)
	}
}

// ErrorHandler serves our error page
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if len(msg) > 0 {
		fmt.Printf("Error: %v", r.URL.Query().Get("msg"))
	} else {
		fmt.Print("No error message specified")
	}
}

// send error function, sets error message and redirects client to error page
func sendError(w http.ResponseWriter, r *http.Request, msg string) {
	http.Redirect(w, r, fmt.Sprintf("/error?msg=%v", msg), 302)
}
