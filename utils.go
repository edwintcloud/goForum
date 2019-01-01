package main

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration is our configuration struct for holding our configurations
type Configuration struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
	Static       string
	Version      string
}

// Initialize our package global variables
var config Configuration
var logger *log.Logger

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
