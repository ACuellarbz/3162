// Filename: cmd/web/users.go
package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()
	//File Server
	fileServer := http.FileServer(http.Dir("./static"))
	router.HandlerFunc(http.MethodGet, "/", app.Home)
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static", fileServer)) //Changed this around to implement the rest capability
	router.HandlerFunc(http.MethodGet, "/tickets", app.Tickets)
	router.HandlerFunc(http.MethodGet, "/schedule", app.Schedule)
	router.HandlerFunc(http.MethodGet, "/login", app.Login)

	return router

}
