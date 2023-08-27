package main

import (
	"log"
	"net/http"
	"flag"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

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

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	//flag.String() returns a pointer to addr, derefernce to access port
	infoLog.Printf("starting server on %s", *addr)

	//Initialize a new http server struct to use custom errorLog instead of
	//standard logger
	srv := &http.Server{
		Addr : *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}