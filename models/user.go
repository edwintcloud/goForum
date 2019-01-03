package models

import (
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

// Session is our session model struct
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
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
