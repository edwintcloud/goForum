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

// ViewThreadHandler serves our view thread page
// GET /threads/read?id=
func ViewThreadHandler(w http.ResponseWriter, r *http.Request) {

	// get id from url query
	id := r.URL.Query().Get("id")
	if len(id) > 0 {

		// create thread struct with data we need
		thread := models.Thread{
			UUID: id,
		}

		// get thread from db
		if err := thread.GetByUUID(); err != nil {
			sendError(w, r, "Unable to find thread")
		} else {

			// get our session
			_, err := session(r)
			if err != nil {
				render(w, &thread, "layout", "public.navbar", "public.thread")
			} else {
				render(w, &thread, "layout", "private.navbar", "private.thread")
			}
		}

	} else {
		sendError(w, r, "Id not specified")
	}
}
