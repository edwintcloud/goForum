package controllers

import (
	"fmt"
	"net/http"
	"time"

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
			utils.Log("error", fmt.Sprintf("Cannot create user - form parse error: %s", err))
		}

		// create user struct instance with values from form
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}

		// create new user in db
		if err := user.Create(); err != nil {
			utils.Log("error", fmt.Sprintf("Unable to create user: %s", err))
			http.Redirect(w, r, "/", 302)
		} else {
			utils.Log("info", fmt.Sprintf("User created: %v", user.Email))
			http.Redirect(w, r, "/login", 302)
		}
	} else {
		sendError(w, r, "Invalid method, POST is required for this endpoint")
	}
}

// LoginHandler serves our login Page
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	render(w, nil, "login.layout", "public.navbar", "login")
}

// AuthenticateHandler logs in a user an sets the session
func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {

	// Only accept POST request, otherwise return error
	if r.Method == http.MethodPost {

		// parse form into r
		err := r.ParseForm()
		if err != nil {
			utils.Log("error", fmt.Sprintf("Cannot login user - unable to parse form: %s", err))
		}

		// create user struct instance with values from form we need
		user := models.User{
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}

		// authenticate user with db
		err = user.Authenticate()
		if err != nil {
			sendError(w, r, fmt.Sprintf("Unable to authenticate user - %s", err))
		} else {

			// create user session in db
			session, err := user.CreateSession()
			if err != nil {
				utils.Log("error", fmt.Sprintf("Cannot login user - unable to create session: %s", err))
				sendError(w, r, "Unable to login due to a server-side issue, please consult an administrator")
			} else {

				// set cookie with session UUID
				cookie := http.Cookie{
					Name:     "session.uuid",
					Value:    session.UUID,
					HttpOnly: true,
					Expires:  time.Now().Add(time.Hour), // Expires in an hour
					Path:     "/",
				}
				http.SetCookie(w, &cookie)

				// redirect to index
				http.Redirect(w, r, "/", 302)
			}
		}
	} else {
		sendError(w, r, "Invalid method, POST is required for this endpoint")
	}
}
