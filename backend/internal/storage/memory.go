package storage

import (
    "contact-hub/backend/internal/model"
    "sync"
)

// PersonStorage (stub) - thread-safe in-memory storage.
type PersonStorage struct {
    mu      sync.RWMutex
    persons []model.Person
}

// NewPersonStorage creates a new instance of the repository.
func NewPersonStorage() *PersonStorage {
    return &PersonStorage{
        persons: make([]model.Person, 0),
    }
}

// TODO: Implement methods for adding, searching, and filtering data.