package main

import (
	"net/http"

)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplates(w, "./static/html/poll.page.tmpl")

}

func (app *application) Tickets(w http.ResponseWriter, r *http.Request) {

}

func (app *application) Schedule(w http.ResponseWriter, r *http.Request) {

}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {

}
