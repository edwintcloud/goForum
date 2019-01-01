package controllers

import (
	"fmt"
	"net/http"
)

// index handler function
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello from %v", r.RequestURI)
	if err != nil {
		fmt.Printf("Unable to complete request: %v\n", err)
	}
}