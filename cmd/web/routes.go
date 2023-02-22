package main

import "net/http"

func (app *application) routes() http.Handler {
	// Swap the route declarations to use the applcation struct's methods as the
	// handler functions.
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Pass the servemux as the 'next' parameter to the secureHeaders middleware
	// Because secureHeaders is just a function, and the function returns a
	// http.Handler we don't need to do anything else.
	// and Wrap the existing chain with the logRequest middleware.
	// and wrap the existing chain with the recoverPanic middleware
	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
