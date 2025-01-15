package main

import (
	"github.com/justinas/alice"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", app.home)                        // Route for the home page
	mux.HandleFunc("/snippet/view", app.snippetView)     // Route to view a specific snippet
	mux.HandleFunc("/snippet/create", app.snippetCreate) // Route to create a new snippet

	//create a middleware chain containing our standard middleware
	//which will be used for every request our application recieves
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	//return app.recoverPanic(app.logRequest(secureHeaders(mux))) // Return the configured multiplexer by wrapping the existing chain
	// with the logRequest middleware which in turn is wrapped under the recoverPanic middleware
	return standard.Then(mux)
}
