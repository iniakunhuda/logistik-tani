package main

import (
	"github.com/gorilla/mux"
)

func (app *application) routes() *mux.Router {
	// Register handler functions.
	r := mux.NewRouter()
	r.HandleFunc("/", app.home)

	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/static/").Handler(app.static("./ui/static/"))
	return r
}
