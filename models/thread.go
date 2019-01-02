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
}

// GetAll gets all threads from the database
func (*Thread) GetAll() ([]*Thread, error) {
	var threads []*Thread

	// Execute query to find all threads
	rows, err := utils.Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}

	// iterate through found rows
	for rows.Next() {
		thread := Thread{}

		// bind row fields to thread struct instance
		if err = rows.Scan(&thread.ID, &thread.UUID, &thread.Topic, &thread.UserID, &thread.CreatedAt); err != nil {
			return nil, err
		}

		// append thread reference to threads slice (memory/speed optimization)
		threads = append(threads, &thread)
	}

	// Close db session pool  and return results
	rows.Close()
	return threads, nil
}
