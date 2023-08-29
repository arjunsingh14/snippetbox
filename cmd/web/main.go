package main

import (
	"database/sql"
	"log"
	"net/http"
	"flag"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"snippetbox.arjun.net/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	snippets *models.SnippetModel
}

func main() {
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime
	// Use "go run ./cmd/web -help" to see all available command line flags
	//Parses the command line flag and assigns it to the addr variable.
	//If there is an error it terminates
	
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")


	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERR\t", log.Ldate|log.Ltime)


	db, err := openDB(*dsn)
	if err != nil{
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &models.SnippetModel{ DB: db },
	}
	//Closes the db connection pool just before the main function exits
	defer db.Close()

	//flag.String() returns a pointer to addr, derefernce to access port
	infoLog.Printf("starting server on %s", *addr)

	//Initialize a new http server struct to use custom errorLog instead of
	//standard logger
	srv := &http.Server{
		Addr : *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}