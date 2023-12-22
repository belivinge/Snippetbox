package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

// http.Handler instead of *http.ServeMux
func (app *application) routes() http.Handler {
	// alice managing middleware
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// the middleware specific to dynamic application routes and nosurf middleware
	dynamicMiddleware := alice.New(app.session.Enable, noSurf)

	// mux := http.NewServeMux()
	mux := pat.New()

	// mux.HandleFunc("/", app.home)
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	// // you can also use instead of alice
	// mux.Get("/", app.session.Enable(http.HandlerFunc(app.home)))

	// snippet form
	mux.Get("/sneep/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.creatorForm))

	// mux.HandleFunc("/sneep", app.snippet)
	mux.Post("/sneep/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.creator))

	// changing the id route URL path
	mux.Get("/sneep/:id", dynamicMiddleware.ThenFunc(app.snippet))

	// User Authentication
	mux.Get("/sneep/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/sneep/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/sneep/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/sneep/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/sneep/user/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(app.logoutUser))

	// static routes, no dynamic applications
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// wrapping the chain with the logrequest middleware
	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return standardMiddleware.Then(mux)
}

// 208 page
