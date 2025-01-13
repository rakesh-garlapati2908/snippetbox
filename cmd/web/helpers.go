package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// Helper function to handle server errors
func (app *application) serverError(w http.ResponseWriter, err error) {
	// Capture the error message along with the stack trace
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	// Log the error with the stack trace
	app.errorLog.Output(2, trace)

	// Send an internal server error response
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Helper function to handle client errors (e.g., invalid requests)
func (app *application) clientError(w http.ResponseWriter, status int) {
	// Send a client error response (e.g., 400 Bad Request, 404 Not Found)
	http.Error(w, http.StatusText(status), status)
}

// Helper function to handle 404 Not Found errors
func (app *application) notFound(w http.ResponseWriter) {
	// Use the clientError function to send a 404 response
	app.clientError(w, http.StatusNotFound)
}
