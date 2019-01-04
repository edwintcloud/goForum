package models

import (
	"time"

	"github.com/edwintcloud/goForum/utils"
)

// Session is our session model struct
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

// CreateSession creates a new session for an existing user
func (u *User) CreateSession() (Session, error) {
	session := Session{}

	// prepare our query for execution on the db
	stmt, err := utils.Db.Prepare(`INSERT INTO sessions (uuid, email, user_id, created_at) 
																 VALUES ($1, $2, $3, $4) RETURNING id, uuid, email, user_id, created_at`)
	if err != nil {
		return session, err
	}

	// defer our query connection pool to close after function has completed
	defer stmt.Close()

	// Execute the query, returning at most one row. Scan result into session
	err = stmt.QueryRow(utils.CreateUUID(), u.Email, u.ID, time.Now()).Scan(&session.ID, &session.UUID, &session.Email, &session.UserID, &session.CreatedAt)
	if err != nil {
		return session, err
	}

	// if all went well, return session and nil
	return session, nil
}

// Get retreives a session from the database
func (s *Session) Get() error {

	// Execute the query, returning at most one row. Scan result into s
	err := utils.Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", s.UUID).
		Scan(&s.ID, &s.UUID, &s.Email, &s.UserID, &s.CreatedAt)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// Delete deletes a session from the database
func (s *Session) Delete() error {

	// prepare our query for execution on the db
	stmt, err := utils.Db.Prepare(`DELETE FROM sessions WHERE uuid = $1`)
	if err != nil {
		return err
	}

	// defer our query connection pool to close after function has completed
	defer stmt.Close()

	// Execute the query with arguments
	_, err = stmt.Exec(s.UUID)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}
