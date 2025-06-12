package handlers

import (
	"encoding/json"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	message := map[string]string{"message": "Hello, world!"}

	json.NewEncoder(w).Encode(message)
}
