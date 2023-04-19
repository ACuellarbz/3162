// Filename: cmd/web/routes.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	//File Server
	fileServer := http.FileServer(http.Dir("./static"))
	dynamicMiddleware := alice.New(app.sessionsManager.LoadAndSave)
	router.Handler(http.MethodGet, "/", dynamicMiddleware.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //Changed this around to implement the rest capability
	router.Handler(http.MethodGet, "/tickets", dynamicMiddleware.ThenFunc(app.tickets))
	router.Handler(http.MethodGet, "/schedule", dynamicMiddleware.ThenFunc(app.schedule))
	router.Handler(http.MethodGet, "/login", dynamicMiddleware.ThenFunc(app.login))
	router.Handler(http.MethodPut, "/schedule", dynamicMiddleware.ThenFunc(app.updateSchedule)) //Update a schedule record or etc
	router.Handler(http.MethodPost, "/schedule", dynamicMiddleware.ThenFunc(app.addRoute))
	router.Handler(http.MethodDelete, "/delete", dynamicMiddleware.ThenFunc(app.deleteR))

	standardMiddleware := alice.New(app.RecoverPanicMiddleware,
		app.logRequestMiddleware,
		securityHeadersMiddleware)

	return standardMiddleware.Then(router)
}
