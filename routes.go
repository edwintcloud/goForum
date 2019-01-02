package main

import (
	"net/http"

	"github.com/edwintcloud/goForum/controllers"
)

func registerRoutes(mux *http.ServeMux) {

	// App routes
	mux.HandleFunc("/", controllers.IndexHandler)
	mux.HandleFunc("/error", controllers.ErrorHandler)

	// Auth routes
	mux.HandleFunc("/users/signup", controllers.SignupHandler)
	mux.HandleFunc("/users", controllers.CreateAccountHandler)

	// View Log file in html format
	mux.HandleFunc("/admin/log", controllers.ViewLogHandler)

}
