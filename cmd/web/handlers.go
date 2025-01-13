package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Home handler for the root URL ("/")
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// If the path is not "/", return a 404 error
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Files to be used in the home page template rendering
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	// Parse the templates
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// If there's an error in parsing, return a server error
		app.serverError(w, err)
		return
	}

	// Execute the template with the "base" layout and no data
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		// If there's an error executing the template, return a server error
		app.serverError(w, err)
		return
	}
}

// Snippet view handler (to view a specific snippet)
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Parse the "id" query parameter to get the snippet ID
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// If the ID is invalid or less than 1, return a 404 error
		app.notFound(w)
		return
	}

	// For now, just output the snippet ID (this will later be replaced with actual snippet retrieval logic)
	fmt.Fprintf(w, "Snippet %d", id)
}

// Snippet create handler (to create a new snippet)
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests for snippet creation
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)        // Set the "Allow" header to indicate allowed methods
		app.clientError(w, http.StatusMethodNotAllowed) // Respond with a 405 Method Not Allowed error
		return
	}
	// For now, just output a simple message
	w.Write([]byte("Create a new snippet!"))
}
