package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home) // subtree path, ends with slash, default - > followed by anything

	mux.HandleFunc("/sneep", snippet)        // fixed path
	mux.HandleFunc("/sneep/create", creator) // fixed path

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting Server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
