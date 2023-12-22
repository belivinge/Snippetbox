package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// it uses a CSRF cookie with the Secure, Path and httpOnly flags set
func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})
	return csrfHandler
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// deferred function which will run as go unwinds the stack
		// goroutine to do some background processing
		defer func() {
			// builtin recover function
			if err := recover(); err != nil {
				// closes the connection and informs the user
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// to help prevent XSS and Clickjacking attacks
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		next.ServeHTTP(w, r)
	})
}

// to add log HTTP requests
// logRequest - > secureHeaders - > servemux - > application Header
func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if the user is not authorized, redirect them to the login page
		if app.authenticatedUser(r) == 0 {
			http.Redirect(w, r, "user/login", 302)
			return
		}
		// otherwise call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
