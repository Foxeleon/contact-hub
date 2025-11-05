package api
import (
	"contact-hub/backend/internal/storage"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Handlers struct holds dependencies like storage.
type Handlers struct {
	Storage *storage.PersonStorage
}

// GetPersons now handles filtering and pagination.
func (h *Handlers) GetPersons(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	query := r.URL.Query().Get("q")
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")
	bdayFromStr := r.URL.Query().Get("birthdayFrom")
	bdayToStr := r.URL.Query().Get("birthdayTo")
	page, _ := strconv.Atoi(pageStr)
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10 // Default page size
	}
	var bdayFrom, bdayTo *time.Time
	if t, err := time.Parse(time.RFC3339, bdayFromStr); err == nil {
		bdayFrom = &t
	}
	if t, err := time.Parse(time.RFC3339, bdayToStr); err == nil {
		bdayTo = &t
	}
	params := storage.SearchParams{
		Query:        query,
		BirthdayFrom: bdayFrom,
		BirthdayTo:   bdayTo,
		Page:         page,
		PageSize:     pageSize,
	}
	result := h.Storage.Query(params)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("ERROR: Failed to encode result to JSON: %v", err)
	}
}