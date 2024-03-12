package main

import (
	"net/http"

	"github.com/justinas/alice"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes(cfg config) http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(cfg.staticDir)})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/curse", app.curse)
	mux.HandleFunc("/view", app.view)
	mux.HandleFunc("/create", app.create)
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(mux)
}
