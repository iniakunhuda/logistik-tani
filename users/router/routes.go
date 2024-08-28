package router

import (
	"github.com/gorilla/mux"
	"github.com/iniakunhuda/logistik-tani/users/controller"
)

func NewRouter(userController *controller.UserController, authController *controller.AuthController, userLandController *controller.UserLandController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", userController.FindAll).Methods("GET")
	r.HandleFunc("/api/users/login", authController.Login).Methods("POST")
	r.HandleFunc("/api/users/profile", authController.Profile).Methods("GET")

	r.HandleFunc("/api/users/{id}", userController.FindById).Methods("GET")
	r.HandleFunc("/api/users", userController.Create).Methods("POST")
	r.HandleFunc("/api/users/{id}", userController.Update).Methods("PUT")
	r.HandleFunc("/api/users/{id}", userController.Delete).Methods("DELETE")

	landRouter := r.PathPrefix("/api/lands").Subrouter()
	landRouter.HandleFunc("", userLandController.FindAll).Methods("GET")
	landRouter.HandleFunc("/{id}", userLandController.FindById).Methods("GET")
	landRouter.HandleFunc("", userLandController.Create).Methods("POST")
	landRouter.HandleFunc("/{id}", userLandController.Update).Methods("PUT")
	landRouter.HandleFunc("/{id}", userLandController.Delete).Methods("DELETE")

	return r
}
