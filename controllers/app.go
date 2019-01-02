package controllers

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/edwintcloud/goForum/models"
)

// IndexHandler serves our index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	thread := models.Thread{}

	// get threads from db
	threads, err := thread.GetAll()
	if err != nil {
		sendError(w, r, "Unable to get threads")
	} else {

		// render template with data and files
		render(w, threads, "layout", "public.navbar", "index")
	}
}

// ErrorHandler serves our error page
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")

	// render template with data and files
	render(w, msg, "layout", "public.navbar", "error")
}

// ViewLogHandler servers our logs as static html
func ViewLogHandler(w http.ResponseWriter, r *http.Request) {
	var log []string

	// Open file for reading
	file, err := os.Open(fmt.Sprintf("%v/goForum.log", curDir()))
	if err != nil {
		sendError(w, r, "Unable to load log file")
	}

	// defer file to close when function exits
	defer file.Close()

	// create new bufio scanner to read through file lines
	scanner := bufio.NewScanner(file)

	// iterate through file lines and append each line to log slice
	for scanner.Scan() {
		log = append(log, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		sendError(w, r, "Unable to read log file lines")
	}

	// render template with data and files
	render(w, log, "layout", "public.navbar", "log")
}

// render function parses template files or returns returns error to client
func render(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string

	// build file list slice
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("%v/views/%v.html", curDir(), file))
	}

	// attempt to parse files
	templates, err := template.ParseFiles(files...)
	if err != nil {
		fmt.Fprintf(w, "Unable to render templates: %v", err)
	} else {

		// render template
		templates.ExecuteTemplate(w, "layout", data)
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
