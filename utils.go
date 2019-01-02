package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Configuration is our configuration struct for holding our configurations
type Configuration struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
	Static       string
	Version      string
	DbHost       string
	DbUser       string
	DbPassword   string
	DbName       string
}

// Initialize our package global variables
var config Configuration
var logger *log.Logger

// Db is our exported db connection instance to be used by the models
var Db *sql.DB

// configuration loader utility function
func loadConfiguration() error {

	// try to open config file
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	// try to decode file into package global config struct using json decoder
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// log loader utility function
func loadLog() error {

	// try to open log file, or create if it doesn't exist
	file, err := os.OpenFile("goForum.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// set package global logger to new logger instance with file
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	// if all went well, return nil
	return nil
}

// connect to database
func connectToDb() error {
	var err error

	// connect to postgres db and set package global Db to Db instance
	Db, err = sql.Open("postgres", fmt.Sprintf(
		"postgres://%v:%v@%s/%s?sslmode=disable",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbName,
	))
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// Initialize database by creating tables if they do not exist
func initializeDb() error {

	// Create users table
	_, err := Db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id         serial primary key,
		uuid       varchar(64) not null unique,
		name       varchar(255),
		email      varchar(255) not null unique,
		password   varchar(255) not null,
		created_at timestamp not null   
	)`)
	if err != nil {
		return err
	}

	// Create sessions table
	_, err = Db.Exec(`CREATE TABLE IF NOT EXISTS sessions (
		id         serial primary key,
		uuid       varchar(64) not null unique,
		email      varchar(255),
		user_id    integer references users(id),
		created_at timestamp not null   
	)`)
	if err != nil {
		return err
	}

	// Create threads table
	_, err = Db.Exec(`CREATE TABLE IF NOT EXISTS threads (
		id         serial primary key,
		uuid       varchar(64) not null unique,
		topic      text,
		user_id    integer references users(id),
		created_at timestamp not null
	)`)
	if err != nil {
		return err
	}

	// Create posts table
	_, err = Db.Exec(`CREATE TABLE IF NOT EXISTS posts (
		id         serial primary key,
		uuid       varchar(64) not null unique,
		body       text,
		user_id    integer references users(id),
		thread_id  integer references threads(id),
		created_at timestamp not null  
	)`)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}
