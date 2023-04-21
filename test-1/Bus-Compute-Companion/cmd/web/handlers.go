package main

import (
	"log"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl", nil)

}

func (app *application) tickets(w http.ResponseWriter, r *http.Request) {

}

func (app *application) scheduleShow(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered Schedule")
	schedule, err := app.bus_schedule.Get()
	if err != nil {
		log.Println(err)
		return
	}
	data := &templateData{
		Schedule: schedule,
	}
	RenderTemplate(w, "schedule.page.tmpl", data)
}
func (app *application) scheduleFormShow(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "schedule.add.tmpl", nil)

}
func (app *application) scheduleFormSubmit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	id := r.PostForm.Get("id")
	company := r.PostForm.Get("company_id")
	begin_location := r.PostForm.Get("begin_id")
	destin_location := r.PostForm.Get("destination_id")
	log.Printf("%s %s %s\n", company, begin_location, destin_location)

	_, err = app.bus_schedule.Insert(id, company, begin_location, destin_location)
	log.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
func (app *application) login(w http.ResponseWriter, r *http.Request) {

}

func (app *application) updateSchedule(w http.ResponseWriter, r *http.Request) {

}

func (app *application) deleteR(w http.ResponseWriter, r *http.Request) {

}