package main

import (
	"log"
	"net/http"
	"flag"
	"os"
)


func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime
	// Use "go run ./cmd/web -help" to see all available command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")

	//Parses the command line flag and assigns it to the addr variable.
	//If there is an error it terminates
	flag.Parse()



	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERR\t", log.Ldate|log.Ltime)

	//create http server multiplexer/request router provides a way to
	//route incoming http  requests to their respective URL path
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	//HandleFunc adds serveHTTP method to passed in handler function thereby
	//making it adhere to the Handler interface and can be used.
	mux.HandleFunc("/", home)
	mux.HandleFunc("snippet/view", snippetView)
	mux.HandleFunc("snippet/create", snippetCreate)

	//flag.String() returns a pointer to addr, derefernce to access port
	infoLog.Printf("starting server on %s", *addr)

	//Initialize a new http server struct to use custom errorLog instead of
	//standard logger
	srv := &http.Server{
		Addr : *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}