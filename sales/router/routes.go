package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/sales/controller"
	"github.com/iniakunhuda/logistik-tani/sales/util"
)

func NewRouter(salesController *controller.SalesController, salesIgmController *controller.SalesIgmController) *mux.Router {
	r := mux.NewRouter()

	sales := r.PathPrefix("/api/sales").Subrouter()
	sales.Use(util.AuthVerify)
	sales.HandleFunc("", salesController.FindAll).Methods("GET")
	sales.HandleFunc("/{id}", salesController.FindById).Methods("GET")
	sales.HandleFunc("", salesController.Create).Methods("POST")
	sales.HandleFunc("/{id}", salesController.Update).Methods("PUT")
	sales.HandleFunc("/{id}", salesController.Delete).Methods("DELETE")

	salesIgm := r.PathPrefix("/api/sales-igm").Subrouter()
	salesIgm.Use(util.AuthVerify)
	salesIgm.HandleFunc("", salesIgmController.FindAll).Methods("GET")
	salesIgm.HandleFunc("/{id}", salesIgmController.FindById).Methods("GET")
	salesIgm.HandleFunc("", salesIgmController.Create).Methods("POST")
	salesIgm.HandleFunc("/{id}", salesIgmController.Update).Methods("PUT")
	salesIgm.HandleFunc("/{id}", salesIgmController.Delete).Methods("DELETE")

	return r
}
