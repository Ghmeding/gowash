package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/internal/models"
)

func (h *Handler) BookTimeSlot(w http.ResponseWriter, r *http.Request) {
	var b models.Booking

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	id, err := h.service.CreateBooking(b.UserID, b.TimeSlotID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(map[string]string{
		"id":      id,
		"message": "Booking created successfully",
	})
}
