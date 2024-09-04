package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/purchase/controller"
	"github.com/iniakunhuda/logistik-tani/purchase/util"
)

func NewRouter(purchaseController *controller.PurchaseController, purchaseIgmController *controller.PurchaseIgmController, purchaseReportsToBankController *controller.PurchaseReportsToBankController) *mux.Router {
	r := mux.NewRouter()

	purchase := r.PathPrefix("/api/purchase").Subrouter()
	purchase.Use(util.AuthVerify)
	purchase.HandleFunc("", purchaseController.FindAll).Methods("GET")
	purchase.HandleFunc("/{id}", purchaseController.FindById).Methods("GET")
	purchase.HandleFunc("", purchaseController.Create).Methods("POST")
	purchase.HandleFunc("/{id}", purchaseController.Update).Methods("PUT")
	purchase.HandleFunc("/{id}", purchaseController.Delete).Methods("DELETE")

	purchaseIgm := r.PathPrefix("/api/purchase-igm").Subrouter()
	// purchaseIgm.Use(util.AuthVerify)
	purchaseIgm.HandleFunc("", purchaseIgmController.FindAll).Methods("GET")
	purchaseIgm.HandleFunc("/{id}", purchaseIgmController.FindById).Methods("GET")
	purchaseIgm.HandleFunc("", purchaseIgmController.Create).Methods("POST")
	purchaseIgm.HandleFunc("/{id}", purchaseIgmController.Delete).Methods("DELETE")

	// setor purchase ke bank
	reportPurchaseIgmToBank := r.PathPrefix("/api/report-to-bank").Subrouter()
	reportPurchaseIgmToBank.HandleFunc("", purchaseReportsToBankController.FindAll).Methods("GET")
	reportPurchaseIgmToBank.HandleFunc("/{id}", purchaseReportsToBankController.FindById).Methods("GET")
	reportPurchaseIgmToBank.HandleFunc("", purchaseReportsToBankController.Create).Methods("POST")
	reportPurchaseIgmToBank.HandleFunc("/{id}", purchaseReportsToBankController.Delete).Methods("DELETE")

	return r
}
