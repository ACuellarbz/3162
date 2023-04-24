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
	fileServer := http.FileServer(http.Dir("./ui/static"))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //Changed this around to implement the rest capability
	dynamicMiddleware := alice.New(app.sessionsManager.LoadAndSave)
	router.Handler(http.MethodGet, "/", dynamicMiddleware.ThenFunc(app.home))
	router.Handler(http.MethodGet, "/tickets", dynamicMiddleware.ThenFunc(app.tickets))
	router.Handler(http.MethodGet, "/login", dynamicMiddleware.ThenFunc(app.login))
	router.Handler(http.MethodGet, "/schedule", dynamicMiddleware.ThenFunc(app.scheduleShow))               //Default Page
	router.Handler(http.MethodGet, "/schedule/create", dynamicMiddleware.ThenFunc(app.scheduleFormShow))    //Show Forums
	router.Handler(http.MethodPost, "/schedule/create", dynamicMiddleware.ThenFunc(app.scheduleFormSubmit)) //Adding Items to the database
	router.Handler(http.MethodGet, "/schedule/update", dynamicMiddleware.ThenFunc(app.updateScheduleShow))            
	router.Handler(http.MethodPost, "/schedule/update", dynamicMiddleware.ThenFunc(app.updateSchedule))             //Update a schedule record or etc
	router.Handler(http.MethodGet, "/schedule/delete", dynamicMiddleware.ThenFunc(app.deleteRouteShow))    //Show Forums
	router.Handler(http.MethodDelete, "/schedule/delete", dynamicMiddleware.ThenFunc(app.deleteRoute))                 //Deleting Records

	standardMiddleware := alice.New(MethodOverride, app.RecoverPanicMiddleware,
		app.logRequestMiddleware,
		securityHeadersMiddleware)

	return standardMiddleware.Then(router)
}
