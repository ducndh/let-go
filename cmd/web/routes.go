package main

import (
	"net/http"

	"github.com/justinas/alice"
	"snippetbox.letgoducndh.net/ui"
)

// The routes() method returns a servemux containing our application routes.
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.FS(ui.Files))
	mux.Handle("GET /static/*filepath", fileServer)

	mux.HandleFunc("GET /ping", ping)

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	mux.Handle("/", dynamic.ThenFunc(app.home))
	mux.Handle("/curse", dynamic.ThenFunc(app.curse))
	mux.Handle("GET /view/{id}", dynamic.ThenFunc(app.view))

	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	protected := dynamic.Append(app.requireAuthentication)

	mux.Handle("GET /create", protected.ThenFunc(app.create))
	mux.Handle("POST /create", protected.ThenFunc(app.createPost))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(mux)
}
