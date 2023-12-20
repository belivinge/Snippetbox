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

	// the middleware specific to dynamic application routes
	dynamicMiddleware := alice.New(app.session.Enable)

	// mux := http.NewServeMux()
	mux := pat.New()

	// mux.HandleFunc("/", app.home)
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	// // you can also use instead of alice
	// mux.Get("/", app.session.Enable(http.HandlerFunc(app.home)))

	// snippet form
	mux.Get("/sneep/create", dynamicMiddleware.ThenFunc(app.creatorForm))

	// mux.HandleFunc("/sneep", app.snippet)
	mux.Post("/sneep/create", dynamicMiddleware.ThenFunc(app.creator))

	// changing the id route URL path
	mux.Get("/sneep/:id", dynamicMiddleware.ThenFunc(app.snippet))

	// User Authentication
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.signupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.signupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.ThenFunc(app.logoutUser))

	// static routes, no dynamic applications
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	// wrapping the chain with the logrequest middleware
	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
	return standardMiddleware.Then(mux)
}

// 208 page
