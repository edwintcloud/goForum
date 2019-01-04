package controllers

import (
	"fmt"
	"net/http"

	"github.com/edwintcloud/goForum/models"
	"github.com/edwintcloud/goForum/utils"
)

// NewThreadHandler serves our create a thread template
// GET /threads/new
func NewThreadHandler(w http.ResponseWriter, r *http.Request) {

	// get our session - this is an authenticated only endpoint
	_, err := session(r)
	if err != nil {
		http.Redirect(w, r, "/users/login", 302)
	} else {
		render(w, nil, "layout", "private.navbar", "new.thread")
	}
}

// CreateThreadHandler creates a new thread
// POST /threads
func CreateThreadHandler(w http.ResponseWriter, r *http.Request) {

	// Only accept POST request, otherwise return error
	if r.Method == http.MethodPost {

		// get our session - this is an authenticated only endpoint
		session, err := session(r)
		if err != nil {
			http.Redirect(w, r, "/users/login", 302)
		} else {

			// parse form into r
			err = r.ParseForm()
			if err != nil {
				utils.Log("error", fmt.Sprintf("Cannot create thread - form parse error: %s", err))
			}

			// create thread struct with values we need to create a thread
			thread := models.Thread{
				Topic:  r.PostFormValue("topic"),
				UserID: session.UserID,
			}

			// create thread
			err = thread.Create()
			if err != nil {
				utils.Log("error", fmt.Sprintf("Cannot create thread in db: %s", err))
			}

			// redirect to index
			http.Redirect(w, r, "/", 302)

		}
	} else {
		sendError(w, r, "Invalid method, POST is required for this endpoint")
	}
}
