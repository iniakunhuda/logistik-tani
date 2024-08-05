package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/sales/controller"
)

func NewRouter(salesController *controller.SalesController) *mux.Router {
	r := mux.NewRouter()

	sales := r.PathPrefix("/api/sales").Subrouter()
	sales.HandleFunc("", salesController.FindAll).Methods("GET")
	sales.HandleFunc("{id}", salesController.FindById).Methods("GET")
	sales.HandleFunc("", salesController.Create).Methods("POST")
	sales.HandleFunc("{id}", salesController.Update).Methods("PUT")
	sales.HandleFunc("{id}", salesController.Delete).Methods("DELETE")

	return r
}
