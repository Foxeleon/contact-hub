package storage

import (
	"contact-hub/backend/internal/model"
	"slices"
	"strings"
	"sync"
	"time"
)

// SearchParams defines the query parameters for searching persons.
type SearchParams struct {
	Query        string
	BirthdayFrom *time.Time
	BirthdayTo   *time.Time
	Page         int
	PageSize     int
}

// QueryResult represents the result of a query, including data and pagination info.
type QueryResult struct {
	Data     []model.Person
	Total    int
	Page     int
	PageSize int
}

// PersonStorage is a thread-safe in-memory storage for Person objects.
type PersonStorage struct {
	mu      sync.RWMutex
	persons []model.Person
}

// NewPersonStorage creates a new instance of PersonStorage.
func NewPersonStorage(initialData []model.Person) *PersonStorage {
	return &PersonStorage{
		persons: initialData,
	}
}

// GetAll returns a copy of all persons from the storage. Leave for testing
func (s *PersonStorage) GetAll() []model.Person {
	s.mu.RLock() // Use a read lock for safety
	defer s.mu.RUnlock()

	// Return a copy to prevent external modification of the internal slice
	personsCopy := make([]model.Person, len(s.persons))
	copy(personsCopy, s.persons)
	return personsCopy
}

// Query applies filtering and pagination to the stored persons.
func (s *PersonStorage) Query(params SearchParams) QueryResult {
	s.mu.RLock()
	defer s.mu.RUnlock()
	filtered := make([]model.Person, 0)
	// 1. Apply full-text search filter (case-insensitive)
	if params.Query != "" {
		searchQuery := strings.ToLower(params.Query)
		for _, p := range s.persons {
			if strings.Contains(strings.ToLower(p.FirstName), searchQuery) ||
				strings.Contains(strings.ToLower(p.LastName), searchQuery) {
				filtered = append(filtered, p)
			}
		}
	} else {
		filtered = append(filtered, s.persons...)
	}

	// 2. Apply date range filter
	dateFiltered := make([]model.Person, 0)
	for _, p := range filtered {
		afterFrom := params.BirthdayFrom == nil || p.Birthday.After(*params.BirthdayFrom) || p.Birthday.Equal(*params.BirthdayFrom)
		beforeTo := params.BirthdayTo == nil || p.Birthday.Before(*params.BirthdayTo) || p.Birthday.Equal(*params.BirthdayTo)
		if afterFrom && beforeTo {
			dateFiltered = append(dateFiltered, p)
		}
	}
	// --- NEW STEP: SORTING ---
	// guarantee stable order before pagination.
	// sort by first name and then by last name for stability.
	slices.SortFunc(dateFiltered, func(a, b model.Person) int {
		if n := strings.Compare(a.FirstName, b.FirstName); n != 0 {
			return n
		}
		return strings.Compare(a.LastName, b.LastName)
	})
	totalRecords := len(dateFiltered)

	page := params.Page
	if page < 1 {
		page = 1 // Default to page 1 if the requested page is invalid
	}
	pageSize := params.PageSize
	if pageSize < 0 {
		pageSize = 10 // Default to a standard page size if negative
	}

	// 3. Apply pagination using sanitized values
	start := (page - 1) * pageSize
	end := start + pageSize
	if start > totalRecords {
		start = totalRecords
	}
	if end > totalRecords {
		end = totalRecords
	}
	paginatedData := dateFiltered[start:end]
	return QueryResult{
		Data:     paginatedData,
		Total:    totalRecords,
		Page:     page, // Return the corrected page number
		PageSize: pageSize,
	}
}
