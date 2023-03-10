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
	dynamicMiddleware := alice.New(app.session.Enable, NoSurf, app.authenticate)

	// Swap the route declarations to use the applcation struct's methods as the
	// handler functions.
	mux := pat.New()

	// Update these routes to use the new dynamicMiddleware chain followed
	// by the appropriate handler function.
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/snippet", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	// Add the requireAuthenticateUser middleware to the chain.
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippetForm))
	// Add the requireAuthenticatedUser middleware to the chain.
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.createSnippet))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippet))

	// Add the five new routes.
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	// Add the requireAuthenticatedUser middleware to the chain.
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))
	mux.Get("/ping", http.HandlerFunc(ping))

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// Return the 'standard' middleware chain followed by the servemux
	return standardMiddleware.Then(mux)
}
