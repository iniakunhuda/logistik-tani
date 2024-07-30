package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/sales/controller"
)

func NewRouter(salesController *controller.SalesController) *mux.Router {
	r := mux.NewRouter()

	bibit := r.PathPrefix("/api/sales/bibit").Subrouter()
	bibit.HandleFunc("", salesController.FindAll).Methods("GET")
	bibit.HandleFunc("/{id}", salesController.FindById).Methods("GET")
	bibit.HandleFunc("", salesController.Create).Methods("POST")
	bibit.HandleFunc("/{id}", salesController.Update).Methods("PUT")
	bibit.HandleFunc("/{id}", salesController.Delete).Methods("DELETE")

	pupuk := r.PathPrefix("/api/sales/pupuk").Subrouter()
	pupuk.HandleFunc("", salesController.FindAll).Methods("GET")
	pupuk.HandleFunc("/{id}", salesController.FindById).Methods("GET")
	pupuk.HandleFunc("", salesController.Create).Methods("POST")
	pupuk.HandleFunc("/{id}", salesController.Update).Methods("PUT")
	pupuk.HandleFunc("/{id}", salesController.Delete).Methods("DELETE")

	return r
}
