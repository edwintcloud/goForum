package models

import (
	"errors"
	"time"

	"github.com/edwintcloud/goForum/utils"
)

// User is our user model struct
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// Create creates a new user
func (u *User) Create() error {

	// prepare our query for execution on the db
	stmt, err := utils.Db.Prepare(`INSERT INTO users (uuid, name, email, password, created_at)
																 VALUES ($1, $2, $3, $4, $5) RETURNING id, uuid, created_at`)
	if err != nil {
		return err
	}

	// defer our query connection pool to close after function has completed
	defer stmt.Close()

	// Execute the query, returning at most one row. Scan result into u
	err = stmt.QueryRow(utils.CreateUUID(), u.Name, u.Email, utils.Encrypt(u.Password), time.Now()).Scan(&u.ID, &u.UUID, &u.CreatedAt)
	if err != nil {
		return err
	}

	// If all went well, return nil
	return nil
}

// Authenticate finds a user by email and ensures given password matches user password
func (u *User) Authenticate() error {

	// find a single user by email and password and bind to user struct
	err := utils.Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1 AND password = $2", u.Email, utils.Encrypt(u.Password)).
		Scan(&u.ID, &u.UUID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		return errors.New("invalid username or password")
	}

	// if all went well, return nil
	return nil
}
