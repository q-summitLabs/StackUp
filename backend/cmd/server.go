package main

import (
	"log"
	"net/http"

	"github.com/q-summitLabs/StackedUp/backend/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/api/hello", handlers.Hello)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	log.Printf("Starting server on port %s...", "8080")
	err := server.ListenAndServe()

	if err != nil {
		log.Printf("err: %s", err.Error())
	}

}
