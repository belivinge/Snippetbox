package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetBox!"))
}

func snippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey! you are using snippet right now"))
}

func creator(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Psst, let's create some snippet duh"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home) // subtree path, ends with slash, default - > followed by anything

	mux.HandleFunc("/sneep", snippet)        // fixed path
	mux.HandleFunc("/sneep/create", creator) // fixed path

	log.Println("Starting Server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
