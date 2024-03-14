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

	dynamic := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("/", dynamic.ThenFunc(app.home))
	mux.Handle("/curse", dynamic.ThenFunc(app.curse))
	mux.Handle("GET /view/{id}", dynamic.ThenFunc(app.view))
	mux.Handle("GET /create", dynamic.ThenFunc(app.create))
	mux.Handle("POST /create", dynamic.ThenFunc(app.createPost))
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(mux)
}
