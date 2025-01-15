package main

import (
	"errors"
	"fmt"
	//"html/template"

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

	data := app.newTemplateData(r)
	data.Snippets = snippets

	//use the new render helper
	app.render(w, http.StatusOK, "home.tmpl", data)
}

// Snippet view handler (to view a specific snippet)
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.tmpl", data)
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
