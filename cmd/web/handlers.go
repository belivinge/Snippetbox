package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/belivinge/Snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// we don't need it because the "/" path matches exactly in Pat package
	// if r.URL.Path != "/" {
	// 	// http.NotFound(w, r)
	// 	app.notFound(w) // using the notfound() helper instead
	// 	return
	// }
	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	// *** use the render helper instead ***
	// // for _, snippet := range s {
	// // 	fmt.Fprintf(w, "%v\n", snippet)
	// // }
	// // slice of snippets
	// data := &templateData{Snippets: s}
	// docs := []string{
	// 	"./ui/html/home_page.html",
	// 	"./ui/html/base_layout.html",
	// 	"./ui/html/footer_partial.html",
	// }
	// ts, err := template.ParseFiles(docs...)
	// if err != nil {
	// 	// log.Println(err.Error)
	// 	// method against application, it can access its fields
	// 	// app.errorLog.Println(err.Error)
	// 	// http.Error(w, "Internal Server Error", 500)
	// 	app.serverError(w, err) // using serverError() helper instead
	// 	return
	// }
	// // pass templateData when executing the template
	// err = ts.Execute(w, data)
	// if err != nil {
	// 	// log.Println(err.Error())
	// 	// method against application
	// 	// app.errorLog.Println(err.Error)
	// 	// http.Error(w, "Internal Server Error", 500)
	// 	app.serverError(w, err) // using serverError() helper instead
	// }
	// // w.Write([]byte("Hello from SnippetBox!"))
	app.render(w, r, "home_page.html", &templateData{
		Snippets: s,
	})
}

// changing the signature of every function here so that it is defined as a method against application
func (app *application) snippet(w http.ResponseWriter, r *http.Request) {
	// Pat doesn't strip the colon from the capture key
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		// http.NotFound(w, r)
		app.notFound(w) // use the notFound() helper instead
		return
	}

	// SnippetModels's GET method to get the data for id. If no math - > Not Found
	s, err := app.snippets.Get(int(id))
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	// *** use the render helper instead ***
	// // struct holding the snippet data
	// data := &templateData{Snippet: s}

	// docs := []string{
	// 	"./ui/html/show_page.html",
	// 	"./ui/html/base_layout.html",
	// 	"./ui/html/footer_partial.html",
	// }
	// // parse the templates
	// ts, err := template.ParseFiles(docs...)
	// if err != nil {
	// 	app.serverError(w, err)
	// 	return
	// }
	// // execute and pass in the templateData
	// err = ts.Execute(w, data)
	// if err != nil {
	// 	app.serverError(w, err)
	// }
	// // w.Write([]byte("Hey! you are using snippet right now"))
	// // fmt.Fprintf(w, "\nDisplay id : %d\n", id)
	// // fmt.Fprintf(w, "%v\n", s)
	app.render(w, r, "show_page.html", &templateData{
		Snippet: s,
	})
}

// returns a placeholder result
func (app *application) creatorForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create_page.html", nil)
}

func (app *application) creator(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// 	w.Header().Set("Allow", "POST")
	// 	// http.Error(w, "Method Not Allowed", 405)
	// 	app.clientError(w, http.StatusMethodNotAllowed) // using the clientError() helper instead
	// 	return
	// }

	// the same way for PUT and PATCH requests
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// some dummy data
	// using PostForm.Get() method to retireve the data from r.PostForm map.
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	// a map to hold any validation errors
	errors := make(map[string]string)

	// checking if title is not empty and is not more than 100 chs long
	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximum is 100 characters)"
	}

	// checking others for validation
	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}
	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This field cannot be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "This field is invalid"
	}

	// if there are any errors
	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

	// pass the data to the snippetmodel method, receiving the id
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// redirect the user to the relevant page
	http.Redirect(w, r, fmt.Sprintf("/sneep/%d", id), http.StatusSeeOther)
	// w.Write([]byte("Psst, let's create some snippet duh"))
}

// func downloadHandler(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./ui/static/file.zip")
// }
