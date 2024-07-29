package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/users/controller"
)

func NewRouter(userController *controller.UserController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/users/", userController.FindAll).Methods("GET")
	r.HandleFunc("/api/users/{id}", userController.FindById).Methods("GET")
	r.HandleFunc("/api/users/", userController.Create).Methods("POST")
	r.HandleFunc("/api/users/{id}", userController.Update).Methods("PUT")
	r.HandleFunc("/api/users/{id}", userController.Delete).Methods("DELETE")

	return r
}
