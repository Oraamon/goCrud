package routes

import (
	"ramori/internal/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(userController *controllers.UserController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/users", userController.Update).Methods("PUT")
	r.HandleFunc("/users", userController.Delete).Methods("DELETE")
	r.HandleFunc("/users/password", userController.UpdatePassword).Methods("PUT")

	return r
}
