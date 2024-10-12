package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"ramori/internal/models"
)

type UserController struct {
	DB *sql.DB
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	id, err := models.CreateUser(uc.DB, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	user.ID = int(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
