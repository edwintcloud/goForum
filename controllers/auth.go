package controllers

import (
	"fmt"
	"net/http"

	"github.com/edwintcloud/goForum/models"
	"github.com/edwintcloud/goForum/utils"
)

// SignupHandler serves our signup page
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	render(w, nil, "login.layout", "public.navbar", "signup")
}

// CreateAccountHandler creates an account in the database and redirects to login
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {

	// Only accept POST requests otherwise return error
	if r.Method == http.MethodPost {

		// parse form into r
		err := r.ParseForm()
		if err != nil {
			utils.Log("error", "Cannot create user - form parse error")
		}

		// create user struct instance with values from form
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}

		// just log for now
		utils.Log("info", fmt.Sprintf("User created: %v", user.Email))

		http.Redirect(w, r, "/login", 302)
	} else {
		sendError(w, r, "Invalid method, POST is required for this endpoint")
	}
}
