package main

import (
	"crypto/tls"
	"database/sql" // new import
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/belivinge/Snippetbox/pkg/models/sqlite"
	"github.com/golangcollege/sessions"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	Addr      string
	StaticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	session  *sessions.Session
	// make SnippetModel object available to handlers
	snippets *sqlite.SnippetModel
	// templatecache field to the app
	templatecache map[string]*template.Template
}

// parsing the runtime configuration settings
// making dependencies for the handlers
// running http server

func main() {
	// var err error
	// addr := flag.String("addr", ":4000", "HTTP network address")
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	// Defining a new command-file flag for MYSQL DSN string
	dsn := flag.String("dsn", "db/snippetbox.db?parseTime=true", "MySQL database")
	// a random key for the session secret
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret key")
	flag.Parse()
	// you can always open a file in Go and use it as your log destination:
	f, err := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// use log.New() to create a logger for writing information messages
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// use stderr for writing error messages
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// to create a connection pool into separate openDB() function - > we pass openDB() the DSN from the flag
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// _, err = db.Exec("ATTACH DATABASE 'snippetbox.db' AS snippetbox")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db, err := sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }
	// fmt.Println("Connected!")
	// defer a call to db.Close(), so that the connection pool is closed before main.go function exists
	defer db.Close()

	// a new template cache
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// to initialize a session manager
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour // expires after 12 hours
	session.Secure = true             // the secure flag is set on our cookies

	// initialize a new instance of application containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		session:  session,
		// adding Snippetbox to the application dependencies
		snippets:      &sqlite.SnippetModel{DB: db},
		templatecache: templateCache,
	}

	// a tls.Config strcut is initialized
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,                                     // controls whether the HTTPS connection should use Go's fovored cipher suites or user's. By setting this to true - we prefer Go's suites.
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256}, // specify which elliptic curves should be preferred during the TLS handshake
		//in Go only CurveP256 and X25519 have assembly implementations
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
		Addr:      cfg.Addr,
		ErrorLog:  errorLog,
		Handler:   app.routes(),
		TLSConfig: tlsConfig,//transport layer security
		// Adding Idle, Read and Write Timeouts to the server
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// writing messages using two loggers, instead of the standard logger
	infoLog.Printf("Starting server on %s", cfg.Addr)
	// $go run cmd/web/* -addr="number more than > 1023"
	// because ports 0-1023 are restricted and can only be used by services which have root privileges

	// err := http.ListenAndServe(*addr, mux)
	// errorLog.Fatal(err)
	// use the listenandserve to start the HTTPS server
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

// The opendb() function wraps sql.Open() and returns a sql.DB connection pool for a given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("aoaoao")
	return db, nil
}
