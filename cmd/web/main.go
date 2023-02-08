package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies for that
// web application. For now we'll only include fields for the tow custom loggers
// we'll add more to ias the build progresses.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

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

	// Use log.New() to create a logger for writing information messages. This function
	// has three parameters: the destinate to write the logs to (os.Stdout), a static
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// Create a logger for writing error messsages in the same way, but use stderr
	// the desitnation and se the log.Lshortfile flag to include the relevante
	// file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Initialize a new http.Server struct. We set the Addr and Handler fields
	// that the server uses the same network address and routes as before, and
	// the ErrorLog field so that the server now uses the custom errorLog logger
	// the event of any problems
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Write messages using the two new loggers, instead of standart logger
	infoLog.Printf("Starting server on %s", *addr)

	// Call the ListenAndServe() method on our new http.Server struct
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
