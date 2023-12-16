package main

import "net/http"

// http.Handler instead of *http.ServeMux
func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/sneep", app.snippet)
	mux.HandleFunc("/sneep/create", app.creator)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return secureHeaders(mux)
}

// 208 page
