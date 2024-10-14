package controllers

import (
	"encoding/json"
	"net/http"
	"ramori/internal/usecases"
	"time"
)

type UserController struct {
	UserUseCase usecases.UserUseCase
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := uc.UserUseCase.CreateUser(input.FirstName, input.LastName, input.Password, input.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
