package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/internal/handlers"
	"example.com/internal/services"
	"example.com/internal/storage"
)

func main() {
	repo := storage.NewMemoryRepo()
	svc := services.NewService(repo)
	h := handlers.NewHandler(svc)
	router := handlers.NewRouter(h)

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("Could not listen on 8080...")
	}
}
