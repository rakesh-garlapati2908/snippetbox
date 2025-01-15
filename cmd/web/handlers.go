package main

import (
	"errors"
	"fmt"
	//"html/template"
	"net/http"
	"snippetbox.rakesh.net/internal/models"
	"strconv"
)

// Home handler for the root URL ("/")
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// If the path is not "/", return a 404 error
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n\n", snippet)
	}

	// Files to be used in the home page template rendering
	//files := []string{
	//	"./ui/html/base.tmpl",
	//	"./ui/html/partials/nav.tmpl",
	//	"./ui/html/pages/home.tmpl",
	//}
	//
	// Parse the templates
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	// If there's an error in parsing, return a server error
	//	app.serverError(w, err)
	//	return
	//}
	//
	//// Execute the template with the "base" layout and no data
	//err = ts.ExecuteTemplate(w, "base", nil)
	//if err != nil {
	//	// If there's an error executing the template, return a server error
	//	app.serverError(w, err)
	//	return
	//}
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
	//Retrieve the Snippet from the Database
	snippet, err := app.snippets.Get(id)
	if err != nil {
		//Checks if the error is models.ErrNoRecord, indicating no snippet with the given ID exists. In this case, a 404 response is sent.
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

// Snippet create handler (to create a new snippet)
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests for snippet creation
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)        // Set the "Allow" header to indicate allowed methods
		app.clientError(w, http.StatusMethodNotAllowed) // Respond with a 405 Method Not Allowed error
		return
	}

	//dummy data
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	//After successfully inserting the snippet, the handler redirects the client to the view page for the new snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
