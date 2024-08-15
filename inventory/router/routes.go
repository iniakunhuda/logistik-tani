package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/inventory/controller"
	"github.com/iniakunhuda/logistik-tani/inventory/util"
)

func NewRouter(inventoryController *controller.InventoryController, inventoryPetaniController *controller.InventoryPetaniController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/inventory/all", inventoryController.FindAllWithoutAuth).Methods("GET")
	r.HandleFunc("/api/inventory/all/detail/{id}", inventoryController.FindByIdWithoutAuth).Methods("GET")
	r.HandleFunc("/api/inventory/all/update_reduce_stock/{id}", inventoryController.UpdateReduceStock).Methods("PUT")
	r.HandleFunc("/api/inventory/all/update_increase_stock/{id}", inventoryController.UpdateIncreaseStock).Methods("PUT")

	inventory := r.PathPrefix("/api/inventory").Subrouter()
	inventory.Use(util.AuthVerify)
	inventory.HandleFunc("", inventoryController.FindAll).Methods("GET")
	inventory.HandleFunc("/{id}", inventoryController.FindById).Methods("GET")
	inventory.HandleFunc("", inventoryController.Create).Methods("POST")
	inventory.HandleFunc("/{id}", inventoryController.Update).Methods("PUT")
	inventory.HandleFunc("/{id}", inventoryController.Delete).Methods("DELETE")

	inventoryPetani := r.PathPrefix("/api/inventory/petani").Subrouter()
	inventoryPetani.HandleFunc("", inventoryPetaniController.Create).Methods("POST")

	return r
}
