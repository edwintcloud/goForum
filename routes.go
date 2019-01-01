package main

import (
	"github.com/edwintcloud/goForum/controllers"
	"net/http"
)

func registerRoutes(mux *http.ServeMux) {

	// Index route
	mux.HandleFunc("/", controllers.IndexHandler)
}
