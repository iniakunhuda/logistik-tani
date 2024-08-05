package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/inventory/controller"
	"github.com/iniakunhuda/logistik-tani/inventory/util"
)

func NewRouter(inventoryController *controller.InventoryController) *mux.Router {
	r := mux.NewRouter()

	inventory := r.PathPrefix("/api/inventory").Subrouter()
	inventory.Use(util.AuthVerify)
	inventory.HandleFunc("", inventoryController.FindAll).Methods("GET")
	inventory.HandleFunc("{id}", inventoryController.FindById).Methods("GET")
	inventory.HandleFunc("", inventoryController.Create).Methods("POST")
	inventory.HandleFunc("{id}", inventoryController.Update).Methods("PUT")
	inventory.HandleFunc("{id}", inventoryController.Delete).Methods("DELETE")

	return r
}
