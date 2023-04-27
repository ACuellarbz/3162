package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var dataStore = struct {
	sync.RWMutex
	data map[string]int64
}{data: make(map[string]int64)}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl", nil)

}

func (app *application) tickets(w http.ResponseWriter, r *http.Request) {

}

// Read Implementation
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

// POST METHOD implementation of Create
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

// Displays Update Request Page
func (app *application) updateScheduleShow(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "schedule.update.request.tmpl", nil)
}

// POST METHOD Implementation for Update
func (app *application) updateSchedule(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	id := r.PostForm.Get("id")
	info, schedule_id, err := app.bus_schedule.SearchRecord(id)

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	schedule, err := strconv.Atoi(schedule_id)
	if err != nil {
		// ... handle error
		panic(err)
	}

	data := &templateData{
		ScheduleByte: info,
	}

	ts, err := template.ParseFiles("./ui/html/schedule.update.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	dataStore.Lock()
	dataStore.data["key"] = int64(schedule)
	dataStore.Unlock()

	log.Println(data)
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}
func (app *application) updateRecords(w http.ResponseWriter, r *http.Request) {
	log.Println("Im inside updateRecords")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	id := r.PostForm.Get("id")
	company := r.PostForm.Get("company_id")
	begin_location := r.PostForm.Get("begin_id")
	destin_location := r.PostForm.Get("destination_id")
	err = app.bus_schedule.Update(id, company, begin_location, destin_location)
	log.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func (app *application) deleteRouteShow(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "schedule.delete.tmpl", nil)

}
func (app *application) deleteRoute(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	id := r.PostForm.Get("id")
	err = app.bus_schedule.Delete(id)

	log.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
