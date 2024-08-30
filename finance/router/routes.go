package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/finance/controller"
	"github.com/iniakunhuda/logistik-tani/finance/util"
)

func NewRouter(payoutController *controller.PayoutHistoryController) *mux.Router {
	r := mux.NewRouter()

	purchase := r.PathPrefix("/api/payout-history").Subrouter()
	purchase.Use(util.AuthVerify)
	purchase.HandleFunc("", payoutController.FindAll).Methods("GET")
	purchase.HandleFunc("/{id}", payoutController.FindById).Methods("GET")
	purchase.HandleFunc("", payoutController.Create).Methods("POST")
	purchase.HandleFunc("/{id}", payoutController.Update).Methods("PUT")
	purchase.HandleFunc("/{id}", payoutController.Delete).Methods("DELETE")

	return r
}
