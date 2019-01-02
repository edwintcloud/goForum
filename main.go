package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/edwintcloud/goForum/utils"
)

// init is executed before main, here we will load configuration and logs
func init() {

	// Load configuration from file using utils.go function
	if err := utils.LoadConfiguration(); err != nil {
		log.Fatalf("Failed to load configuration file: %v\n", err)
	}

	// Load logs from file using utils.go function
	if err := utils.LoadLog(); err != nil {
		log.Fatalf("Failed to load log file: %v\n", err)
	}

	// Connect to database using utils.go function
	if err := utils.ConnectToDb(); err != nil {
		log.Fatalf("Failed to connect to db: %v\n", err)
	}

	// Initialize database using utils.go function
	if err := utils.InitializeDb(); err != nil {
		log.Fatalf("Failed to create neccessary tables for db: %v\n", err)
	}

}

// main entry point of our program
func main() {

	// Print server started status to stdout
	fmt.Printf("goForum %v started at %v", utils.Config.Version, utils.Config.Address)

	// Create new http request multiplexer and setup to server static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(utils.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// Register routes with the multiplexer, routes are in the routes.go file
	registerRoutes(mux)

	// Configure settings for http server
	server := &http.Server{
		Addr:           utils.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(utils.Config.ReadTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(utils.Config.WriteTimeout * int(time.Second)),
		MaxHeaderBytes: 1 << 20, // 1 bit shifted left by 20 = 1048576 in decimal
	}

	// Start http server
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Unable to start http server: %v\n", err)
	}
}
