package handlers

import (
	"net/http"
)

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /user", h.SignUp)
	mux.HandleFunc("POST /booking", h.BookTimeSlot)

	return mux
}
