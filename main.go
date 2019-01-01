package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// init is executed before main, here we will load configuration and logs
func init() {
	// loadConfiguration() and loadLog() are found in utils.go
	// both return error or nil if no error
	// If there is an error, program will exit with code 1
	if err := loadConfiguration(); err != nil {
		log.Fatalf("Failed to load configuration file: %v\n", err)
	}
	if err := loadLog(); err != nil {
		log.Fatalf("Failed to load log file: %v\n", err)
	}
}

// main entry point of our program
func main() {

	// Print server started status to stdout
	fmt.Printf("goForum %v started at %v", config.Version, config.Address)

	// Create new http request multiplexer and setup to server static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Register routes with the multiplexer, routes are in the routes.go file
	registerRoutes(mux)

	// Configure settings for http server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20, // 1 bit shifted left by 20 = 1048576 in decimal
	}

	// Start http server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Unable to start http server: %v\n", err)
	}
}
