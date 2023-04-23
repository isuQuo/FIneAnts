package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// Create a new httprouter instance.
	router := httprouter.New()

	// Wrap the http.NotFound() function in a http.HandlerFunc so that it
	// returns our own custom 404 Not Found response.
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})

	dynamic := alice.New(
		app.sessionManager.LoadAndSave,
		// Disable if no TLS
		//app.noSurf,
		app.authenticate,
	)

	// Register the relevant methods, URLs and handlers for the dynamic routes.
	//router.Handler(http.MethodGet, "/user/signup", dynamic.ThenFunc(app.signupUserForm))

	protected := dynamic.Append(app.requireAuthentication)

	router.Handler(http.MethodGet, "/", protected.ThenFunc(app.home))

	// Create a middleware chain containing our 'standard' middleware
	// which will be used by every request our application receives.
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
