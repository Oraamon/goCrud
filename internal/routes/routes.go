package routes

import (
	"ramori/internal/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(userController *controllers.UserController) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")

	return r
}
