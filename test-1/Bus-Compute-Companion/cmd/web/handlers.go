package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl", nil)

}

func (app *application) tickets(w http.ResponseWriter, r *http.Request) {

}

func (app *application) schedule(w http.ResponseWriter, r *http.Request) {

}

func (app *application) login(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateSchedule(w http.ResponseWriter, r *http.Request) {

}
func (app *application) addRoute(w http.ResponseWriter, r *http.Request) {

}
func (app *application) deleteR(w http.ResponseWriter, r *http.Request) {

}
