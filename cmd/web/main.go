package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// type Config struct {
// 	Addr      string
// 	StaticDir string
// }

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

// parsing the runtime configuration settings
// making dependencies for the handlers
// running http server

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// cfg := new(Config)
	// flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	// flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	// you can always open a file in Go and use it as your log destination:
	// f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// use log.New() to create a logger for writing information messages
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// use stderr for writing error messages
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize a new instance of application containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", app.home)                // subtree path, ends with slash, default - > followed by anything
	// mux.HandleFunc("/sneep", app.snippet)        // fixed path
	// mux.HandleFunc("/sneep/create", app.creator) // fixed path

	// fileServer := http.FileServer(http.Dir("./ui/static/"))
	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// a new http.Server struct. We set the Addr and Handler fields
	// the ErrorLog field so that the server uses the custom erroring logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// writing messages using two loggers, instead of the standard logger
	infoLog.Printf("Starting server on %s", *addr)
	// $go run cmd/web/* -addr="number more than > 1023"
	// because ports 0-1023 are restricted and can only be used by services which have root privileges

	// err := http.ListenAndServe(*addr, mux)
	// errorLog.Fatal(err)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
