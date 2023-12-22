package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/justinas/nosurf"
)

// takes a pointer to a templateData struct, adds the current year, then returns the pointer
func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CSRFToken = nosurf.Token(r)
	td.AuthenticatedUser = app.authenticatedUser(r)
	td.CurrentYear = time.Now().Year()
	// add the flash message if one exists
	td.Flash = app.session.PopString(r, "flash")
	return td
}

// retirieve the template set from the cache base on html page
func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templatecache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}
	// a new buffer
	buf := new(bytes.Buffer)
	// execute the template set passing in dynamic data
	// write the template to the buffer
	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}
	// we pass http.ResponseWriter to a function that takes an io.Writer
	buf.WriteTo(w)
}

// 500 Internal Server Error
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// sends a specific status code and corresponding description. For example: 400 Bad request
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// 404 Not Found
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) authenticatedUser(r *http.Request) int {
	return app.session.GetInt(r, "userID")
}
