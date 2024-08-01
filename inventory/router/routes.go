package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/inventory/controller"
)

func NewRouter(inventoryController *controller.InventoryController) *mux.Router {
	r := mux.NewRouter()

	bibit := r.PathPrefix("/api/inventory/stok/bibit").Subrouter()
	bibit.HandleFunc("", inventoryController.FindAll).Methods("GET")
	bibit.HandleFunc("/{id}", inventoryController.FindById).Methods("GET")
	bibit.HandleFunc("", inventoryController.Create).Methods("POST")
	bibit.HandleFunc("/{id}", inventoryController.Update).Methods("PUT")
	bibit.HandleFunc("/{id}", inventoryController.Delete).Methods("DELETE")

	pupuk := r.PathPrefix("/api/inventory/stok/pupuk").Subrouter()
	pupuk.HandleFunc("", inventoryController.FindAll).Methods("GET")
	pupuk.HandleFunc("/{id}", inventoryController.FindById).Methods("GET")
	pupuk.HandleFunc("", inventoryController.Create).Methods("POST")
	pupuk.HandleFunc("/{id}", inventoryController.Update).Methods("PUT")
	pupuk.HandleFunc("/{id}", inventoryController.Delete).Methods("DELETE")

	return r
}
