package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new command line flag with the name "addr", a default value of ":4000"
	// and some short help that explaining what the flag controls. The value of
	// flag will be stored in the addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use te flag.Parse() function to parse the command line
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000"
	// If any error encountered during parsing te application will be terminated
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// The value returned from te flag.String() functions is a pointer the flag
	// value, not the value itself. So we need to dereference the pointer
	// before using it
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
