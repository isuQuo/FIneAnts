package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
)

type application struct {
	// errorLog and infoLog are used to log messages to the console.
	errorLog *log.Logger
	infoLog  *log.Logger
	// templateCache is a map of parsed HTML templates.
	templateCache map[string]*template.Template
	// formDecoder is used to decode form data.
	formDecoder *form.Decoder
	// sessionManager is used to manage user sessions.
	sessionManager *scs.SessionManager
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Can be passed as a flag
	db, err := openDB("PLACE DSN HERE")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	// Create a new session manager
	sessionManager := scs.New()
	// TODO: Set the session store to use your database
	//sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	// Create a TLS config struct
	/* tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256}, // Prefer X25519, then CurveP256
	} */

	srv := &http.Server{
		Addr:     "127.0.0.1:8080",
		ErrorLog: errorLog,
		Handler:  app.routes(),
		//TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,      // IdleTimeout is the maximum amount of time to wait for the next request when keep-alives are enabled.
		ReadTimeout:  5 * time.Second,  // ReadTimeout is the maximum duration for reading the entire request, including the body.
		WriteTimeout: 10 * time.Second, // WriteTimeout is the maximum duration before timing out writes of the response.
	}

	infoLog.Println("Starting server on :8080")
	err = srv.ListenAndServe()
	//err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	// Open a new connection to the database.
	// TODO: Change the driver name to match the database you are using.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
