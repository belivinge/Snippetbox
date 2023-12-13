package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		app.notFound(w) // using the notfound() helper instead
		return
	}
	docs := []string{
		"./ui/html/home_page.html",
		"./ui/html/base_layout.html",
		"./ui/html/footer_partial.html",
	}
	ts, err := template.ParseFiles(docs...)
	if err != nil {
		// log.Println(err.Error)
		// method against application, it can access its fields
		// app.errorLog.Println(err.Error)
		// http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err) // using serverError() helper instead
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		// log.Println(err.Error())
		// method against application
		// app.errorLog.Println(err.Error)
		// http.Error(w, "Internal Server Error", 500)
		app.serverError(w, err) // using serverError() helper instead
	}
	// w.Write([]byte("Hello from SnippetBox!"))
}

// changing the signature of every function here so that it is defined as a method against application
func (app *application) snippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.notFound(w) // use the notFound() helper instead
		return
	}
	w.Write([]byte("Hey! you are using snippet right now"))
	fmt.Fprintf(w, "\nDisplay id : %d", id)
}

func (app *application) creator(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		// http.Error(w, "Method Not Allowed", 405)
		app.clientError(w, http.StatusMethodNotAllowed) // using the clientError() helper instead
		return
	}

	// some dummy data
	title := "O snail"
	content := "O snail\nClimb Mount Fuje\nBut slowly, slowly\n"
	expires := "7"

	// pass the data to the snippetmodel method, receiving the id
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// redirect the user to the relevant page
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
	w.Write([]byte("Psst, let's create some snippet duh"))
}

// func downloadHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./ui/static/file.zip")
// }
