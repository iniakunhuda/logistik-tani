package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/inventory/controller"
	"github.com/iniakunhuda/logistik-tani/inventory/util"
)

func NewRouter(productController *controller.ProductController, productPetaniController *controller.ProductPetaniController, productionController *controller.ProductionController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/inventory/all", productController.FindAllWithoutAuth).Methods("GET")
	r.HandleFunc("/api/inventory/all/detail/{id}", productController.FindByIdWithoutAuth).Methods("GET")
	r.HandleFunc("/api/inventory/all/update_reduce_stock/{id}", productController.UpdateReduceStock).Methods("PUT")
	r.HandleFunc("/api/inventory/all/update_increase_stock/{id}", productController.UpdateIncreaseStock).Methods("PUT")

	inventory := r.PathPrefix("/api/inventory").Subrouter()
	inventory.Use(util.AuthVerify)
	inventory.HandleFunc("", productController.FindAll).Methods("GET")
	inventory.HandleFunc("/{id}", productController.FindById).Methods("GET")
	inventory.HandleFunc("", productController.Create).Methods("POST")
	inventory.HandleFunc("/{id}", productController.Update).Methods("PUT")
	inventory.HandleFunc("/{id}", productController.Delete).Methods("DELETE")

	inventoryPetani := r.PathPrefix("/api/inventory/petani").Subrouter()
	inventoryPetani.HandleFunc("", productPetaniController.Create).Methods("POST")

	// ========================
	// Panen Petani API
	// ========================

	r.HandleFunc("/api/panen/all", productionController.FindAllWithoutAuth).Methods("GET")
	r.HandleFunc("/api/inventory/all/detail/{id}", productionController.FindByIdWithoutAuth).Methods("GET")

	production := r.PathPrefix("/api/panen").Subrouter()
	production.Use(util.AuthVerify)
	production.HandleFunc("", productionController.FindAll).Methods("GET")
	production.HandleFunc("/history", productionController.CreateRiwayat).Methods("POST")
	production.HandleFunc("/{id}", productionController.FindById).Methods("GET")
	production.HandleFunc("", productionController.Create).Methods("POST")
	production.HandleFunc("/{id}", productionController.Update).Methods("PUT")

	return r
}
