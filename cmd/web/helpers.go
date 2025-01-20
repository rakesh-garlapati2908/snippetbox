package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-playground/form/v4"
	"github.com/justinas/nosurf"
	"net/http"
	"runtime/debug"
	"time"
)

// Helper function to handle server errors
func (app *application) serverError(w http.ResponseWriter, err error) {
	// Capture the error message along with the stack trace
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	// Log the error with the stack trace
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Helper function to handle client errors (e.g., invalid requests)
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Helper function to handle 404 Not Found errors
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {

	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("page '%s' not found in template cache", page)
		app.serverError(w, err)
		return
	}

	//initialize a new buffer
	buf := new(bytes.Buffer)

	//execute the template set and write the response body
	// earlier instead of http.ResponseWriter here we use buf
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

	// write out the provided HTTP status code
	w.WriteHeader(status)

	//pass the contents of the buffer to the http.ResponseWriter
	buf.WriteTo(w)
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
		//add the flash message to the template data, if one exists
		Flash:           app.sessionManager.PopString(r.Context(), "flash"),
		IsAuthenticated: app.isAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	}
}

func (app *application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError
		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}
		return err
	}
	return nil
}

func (app *application) isAuthenticated(r *http.Request) bool {
	return app.sessionManager.Exists(r.Context(), "authenticatedUserID")
}
