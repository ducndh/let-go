package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/alexedwards/scs/pgxstore"
	"github.com/alexedwards/scs/v2"

	"github.com/go-playground/form/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"snippetbox.letgoducndh.net/internal/models"
)

type config struct {
	addr      string
	dsn       string
	staticDir string
}

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	var cfg config
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.dsn, "dsn", "dbname=mydb port=5432 sslmode=disable", "Postgres data source name")
	flag.StringVar(&cfg.staticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dbpool, err := pgxpool.New(context.Background(), cfg.dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = pgxstore.New(dbpool)
	sessionManager.Lifetime = 12 * time.Hour

	sessionManager.Cookie.Secure = true

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		snippets:       &models.SnippetModel{DB: dbpool},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:         cfg.addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(cfg),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", cfg.addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}
