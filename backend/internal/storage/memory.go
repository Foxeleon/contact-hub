package storage

import (
	"contact-hub/backend/internal/model"
	"sync"
)

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

// GetAll returns a copy of all persons from the storage.
func (s *PersonStorage) GetAll() []model.Person {
	s.mu.RLock() // Use a read lock for safety
	defer s.mu.RUnlock()

	// Return a copy to prevent external modification of the internal slice
	personsCopy := make([]model.Person, len(s.persons))
	copy(personsCopy, s.persons)
	return personsCopy
}