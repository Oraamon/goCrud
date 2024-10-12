package controllers

import (
	"encoding/json"
	"net/http"
	"ramori/internal/repositories"
)

type UserController struct {
	Repo repositories.UserRepository
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user repositories.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := uc.Repo.Create(user)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	user.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
