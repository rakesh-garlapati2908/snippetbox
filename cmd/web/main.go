package main

import (
	"database/sql"
	"flag"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"os"
	"snippetbox.rakesh.net/internal/models"
)

// application struct holds loggers for error and info logging.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
}

func main() {
	addr := flag.String("addr", ":4000", "http service address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL datasource name")

	flag.Parse()

	// Create loggers for info and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	//initialize a new template cache...
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	//initialize a decoder instance
	formDecoder := form.NewDecoder()

	// Initialize the application with the loggers
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      &models.SnippetModel{DB: db},
		templateCache: templateCache,
		formDecoder:   formDecoder,
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
	err = srv.ListenAndServe()
	errorLog.Fatal(err) // If there's an error starting the server, log it and exit.
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
