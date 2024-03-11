package main

import (
	"net/http"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes(cfg config) *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(cfg.staticDir)})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/curse", app.Curse)
	mux.HandleFunc("/view", app.View)
	mux.HandleFunc("/create", app.Create)
	return mux
}
