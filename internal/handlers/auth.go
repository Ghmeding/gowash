package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/internal/models"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateUser(u.Email, u.Pwd, u.Username, u.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{
		"id":      id,
		"message": "User created successfully",
	})
}
