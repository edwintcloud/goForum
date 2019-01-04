package models

import (
	"time"

	"github.com/edwintcloud/goForum/utils"
)

// Thread is our thread struct model
type Thread struct {
	ID        int
	UUID      string
	Topic     string
	UserID    int
	CreatedAt time.Time
	CreatedBy string
}

// GetAll gets all threads from the database
func (*Thread) GetAll() ([]*Thread, error) {
	var threads []*Thread

	// Execute query to find all threads
	rows, err := utils.Db.Query("SELECT threads.id, threads.uuid, threads.topic, threads.user_id, threads.created_at, users.name AS created_by FROM threads INNER JOIN users ON threads.user_id = users.id ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}

	// iterate through found rows
	for rows.Next() {
		thread := Thread{}

		// bind row fields to thread struct instance
		if err = rows.Scan(&thread.ID, &thread.UUID, &thread.Topic, &thread.UserID, &thread.CreatedAt, &thread.CreatedBy); err != nil {
			return nil, err
		}

		// append thread reference to threads slice (memory/speed optimization)
		threads = append(threads, &thread)
	}

	// Close db session pool  and return results
	rows.Close()
	return threads, nil
}

// Create creates a new thread
func (t *Thread) Create() error {

	// prepare statement for execution on the db
	stmt, err := utils.Db.Prepare(`INSERT INTO threads (uuid, topic, user_id, created_at) 
										VALUES ($1, $2, $3, $4) RETURNING id, uuid, topic, user_id, created_at`)
	if err != nil {
		return err
	}

	// defer our query connection pool to close after function has completed
	defer stmt.Close()

	// Execute the query, returning at most one row. Scan result into t
	err = stmt.QueryRow(utils.CreateUUID(), t.Topic, t.UserID, time.Now()).Scan(&t.ID, &t.UUID, &t.Topic, &t.UserID, &t.CreatedAt)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}
