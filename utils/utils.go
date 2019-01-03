package utils

import (
	"crypto/rand"
	"crypto/sha1"
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

// Config is our exported configuration instance
var Config Configuration

// Logger is our exported logger instance
var Logger *log.Logger

// Db is our exported db connection instance to be used by the models
var Db *sql.DB

// LoadConfiguration is a configuration loader utility function
func LoadConfiguration() error {

	// try to open config file
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	// try to decode file into package global config struct using json decoder
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Config)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// LoadLog is a log loader utility function
func LoadLog() error {

	// try to open log file, or create if it doesn't exist
	file, err := os.OpenFile("goForum.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	// set package global logger to new logger instance with file
	Logger = log.New(file, "INFO ", log.Ldate|log.Ltime)

	// if all went well, return nil
	return nil
}

// ConnectToDb connects to database
func ConnectToDb() error {
	var err error

	// connect to postgres db and set package global Db to Db instance
	Db, err = sql.Open("postgres", fmt.Sprintf(
		"postgres://%v:%v@%s/%s?sslmode=disable",
		Config.DbUser,
		Config.DbPassword,
		Config.DbHost,
		Config.DbName,
	))
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}

// InitializeDb creates tables if they do not exist
func InitializeDb() error {

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

// Log appends to the log
func Log(t string, args ...interface{}) {
	switch t {
	case "info":
		Logger.SetPrefix("INFO ")
	case "error":
		Logger.SetPrefix("ERROR ")
	case "warning":
		Logger.SetPrefix("WARNING ")
	default:
		Logger.SetPrefix("INFO ")
	}
	Logger.Println(args...)
}

// CreateUUID generated a random UUID from RFC 4122
func CreateUUID() string {

	// allocate a 16 length byte slice
	u := make([]byte, 16)

	// generate 16 cryptographically secure pseudorandom numbers from rand.Reader
	// and write them into our byte slice
	_, err := rand.Read(u)
	if err != nil {
		Log("error", fmt.Sprintf("Unable to generate UUID: %s", err))
	}

	// 0x40 is reserved variant from RFC 4122
	// (u[8] or 64) and 127, ensures u[8] is not larger than 127
	u[8] = (u[8] | 0x40) & 0x7F

	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	// (u[6] and 15) or (64), ensures u[6] is not larger than 15 and adds 64
	u[6] = (u[6] & 0xF) | (0x4 << 4)

	// return formatted uuid, %x translates to base 16, lower-case, two characters per byte
	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

// Encrypt returns a SHA-1 hash from plain text
func Encrypt(t string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(t)))
}
