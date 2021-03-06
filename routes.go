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
	mux.HandleFunc("/users/login", controllers.LoginHandler)
	mux.HandleFunc("/users/authenticate", controllers.AuthenticateHandler)
	mux.HandleFunc("/users/logout", controllers.LogoutHandler)

	// View Log file in html format
	mux.HandleFunc("/admin/log", controllers.ViewLogHandler)

	// Thread routes
	mux.HandleFunc("/threads/new", controllers.NewThreadHandler)
	mux.HandleFunc("/threads", controllers.CreateThreadHandler)
	mux.HandleFunc("/threads/read", controllers.ViewThreadHandler)

}
