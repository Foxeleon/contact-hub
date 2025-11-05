package api

import (
	"contact-hub/backend/internal/storage"
	"encoding/json"
	"log"
	"net/http"
)

// Handlers struct holds dependencies like storage.
type Handlers struct {
	Storage *storage.PersonStorage
}

// GetPersons handles the request to retrieve all person records.
func (h *Handlers) GetPersons(w http.ResponseWriter, r *http.Request) {
	// 1. Get all persons from the storage
	persons := h.Storage.GetAll()

	// 2. For now, return the full list. Pagination and filtering will be added later.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(persons); err != nil {
		log.Printf("ERROR: Failed to encode persons to JSON: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}