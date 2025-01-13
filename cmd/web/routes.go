package main

import "net/http"

// Function to define application routes
func (app *application) routes() *http.ServeMux {
	// Create a new ServeMux (HTTP request multiplexer)
	mux := http.NewServeMux()

	// Set up a file server for static assets like CSS, JS, images
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer)) // Serve static files under "/static/"

	// Define routes and the corresponding handler functions
	mux.HandleFunc("/", app.home)                        // Route for the home page
	mux.HandleFunc("/snippet/view", app.snippetView)     // Route to view a specific snippet
	mux.HandleFunc("/snippet/create", app.snippetCreate) // Route to create a new snippet

	return mux // Return the configured multiplexer
}
