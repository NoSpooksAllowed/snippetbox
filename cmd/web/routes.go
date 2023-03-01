package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// Create a middleware chain containing our 'standard' middleware
	// which will be used for every request our application receives.
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Create a new middleware chain containing the middleware specific to our
	// dynamic application routes. For now, this chain will only contain
	// the session middleware but we'll add more to it later.
	dynamicMiddleware := alice.New(app.session.Enable)

	// Swap the route declarations to use the applcation struct's methods as the
	// handler functions.
	mux := pat.New()

	// Update these routes to use the new dynamicMiddleware chain followed
	// by the appropriate handler function.
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// Return the 'standard' middleware chain followed by the servemux
	return standardMiddleware.Then(mux)
}
