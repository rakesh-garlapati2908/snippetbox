package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application struct holds loggers for error and info logging.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Setting up the address for the HTTP server to listen on, default is ":4000"
	addr := flag.String("addr", ":4000", "http service address")
	flag.Parse()

	// Create loggers for info and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	// Initialize the application with the loggers
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Create a new HTTP server with specific address, error log, and handler (routes)
	srv := &http.Server{
		Addr:     *addr,        // The address to listen on (default ":4000")
		ErrorLog: errorLog,     // Error log for server
		Handler:  app.routes(), // The HTTP request handler (defined in routes.go)
	}

	// Log server startup message
	infoLog.Printf("Starting server on %s", *addr)

	// Start the HTTP server, listen for requests, and log fatal errors
	err := srv.ListenAndServe()
	errorLog.Fatal(err) // If there's an error starting the server, log it and exit.
}
