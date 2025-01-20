package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func secureHeaders(next http.Handler) http.Handler {
	// Return an HTTP handler function that wraps the next handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Security-Policy (CSP) header.
		// This controls the sources from which content like scripts, styles, and fonts can be loaded.
		// In this case:
		// - `default-src 'self'`: Only allow content from the same origin.
		// - `style-src 'self' fonts.googleapis.com`: Allow styles from the same origin and Google Fonts.
		// - `font-src fonts.gstatic.com`: Allow fonts only from Google's font CDN.
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")

		// Set the Referrer-Policy header.
		// This restricts what information is sent in the Referer header when navigating away from your site.
		// `origin-when-cross-origin`: Send the full origin for same-origin requests but only the origin for cross-origin requests.
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")

		// Set the X-Content-Type-Options header.
		// This prevents browsers from trying to guess the MIME type of a file and ensures it is handled as declared.
		// Value `nosniff` prevents MIME-sniffing attacks.
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Set the X-Frame-Options header.
		// This prevents the page from being embedded in an `<iframe>` by other websites.
		// Value `deny` completely disallows embedding, protecting against clickjacking attacks.
		w.Header().Set("X-Frame-Options", "deny")

		// Set the X-XSS-Protection header.
		// This disables some outdated browser-based XSS protection mechanisms.
		// Value `0` turns off the XSS auditor, which can sometimes introduce vulnerabilities in modern contexts.
		w.Header().Set("X-XSS-Protection", "0")

		// Call the next handler in the chain, passing along the response writer and request.
		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	// Return an HTTP handler function that wraps the next handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the details of the incoming request using the application's info logger.
		// The log format includes:
		// - `r.RemoteAddr`: The client's IP address.
		// - `r.Proto`: The protocol used (e.g., HTTP/1.1).
		// - `r.Method`: The HTTP method (e.g., GET, POST).
		// - `r.URL.RequestURI()`: The requested URI (path and query string).
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		// Call the next handler in the middleware chain, passing along the response writer and request.
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use `defer` to ensure this function runs at the end of the handler, even if a panic occurs.
		defer func() {
			// Check if a panic occurred by recovering it.
			if err := recover(); err != nil {
				// If a panic occurred, set the "Connection" header to "close".
				// This signals the client that the server will close the connection after the response.
				w.Header().Set("Connection", "close")

				// Log the panic as a server error for debugging purposes.
				// The panic value (`err`) is converted to an `error` type using `fmt.Errorf`.
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		// Call the next handler in the chain.
		// If this handler causes a panic, the `defer` block will handle it.
		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAuthenticated(r) {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}

		w.Header().Add("Cache-Control", "no-cache")

		next.ServeHTTP(w, r)
	})
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})
	return csrfHandler
}
