package main

import (
	"net/http"

	"github.com/edwintcloud/goForum/controllers"
)

func registerRoutes(mux *http.ServeMux) {

	// App routes
	mux.HandleFunc("/", controllers.IndexHandler)
	mux.HandleFunc("/error", controllers.ErrorHandler)

}
