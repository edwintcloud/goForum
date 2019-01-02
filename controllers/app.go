package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// IndexHandler serves our index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	// parse template files to ensure they are valid
	templates := template.Must(template.ParseFiles(
		fmt.Sprintf("%v/views/layout.html", curDir()),
		fmt.Sprintf("%v/views/public.navbar.html", curDir()),
		fmt.Sprintf("%v/views/index.html", curDir()),
	))

	// respond with layout template
	templates.ExecuteTemplate(w, "layout", "")

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

// get working directory function to return the current working directory
func curDir() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error while getting current directory: %v\n", err)
	}
	return wd
}
