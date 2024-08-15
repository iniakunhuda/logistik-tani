package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/purchase/controller"
	"github.com/iniakunhuda/logistik-tani/purchase/util"
)

func NewRouter(purchaseController *controller.PurchaseController) *mux.Router {
	r := mux.NewRouter()

	purchase := r.PathPrefix("/api/purchase").Subrouter()
	purchase.Use(util.AuthVerify)
	purchase.HandleFunc("", purchaseController.FindAll).Methods("GET")
	purchase.HandleFunc("/{id}", purchaseController.FindById).Methods("GET")
	purchase.HandleFunc("", purchaseController.Create).Methods("POST")
	purchase.HandleFunc("/{id}", purchaseController.Update).Methods("PUT")
	purchase.HandleFunc("/{id}", purchaseController.Delete).Methods("DELETE")

	return r
}
