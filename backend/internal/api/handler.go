package api

import (
	"contact-hub/backend/internal/storage"
	"encoding/json"
	"log"
	"net/http"
)

// Handlers struct to hold dependencies like storage.
type Handlers struct {
	Storage *storage.PersonStorage
}

// GetPersons handles the request to retrieve all person records.
func (h *Handlers) GetPersons(w http.ResponseWriter, r *http.Request) {
	// This is a placeholder. I will fetch real data in the next steps.
	// For now, it returns an empty list, which is a valid response.
	persons := make([]map[string]string, 0)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(persons); err != nil {
		log.Printf("ERROR: Failed to encode persons to JSON: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}