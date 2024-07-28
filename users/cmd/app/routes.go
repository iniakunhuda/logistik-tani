package main

import (
	"github.com/gorilla/mux"
)

func (app *Application) Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/users/", app.all).Methods("GET")
	r.HandleFunc("/api/users/{id}", app.findByID).Methods("GET")
	r.HandleFunc("/api/users/", app.insert).Methods("POST")
	r.HandleFunc("/api/users/{id}", app.update).Methods("PUT")

	return r
}
