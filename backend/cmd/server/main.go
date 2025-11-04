package main

import (
	"contact-hub/backend/internal/api"
	"contact-hub/backend/internal/storage"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
    log.Println("Starting Contact Hub service...")

    // Create storage instance
    personStorage := storage.NewPersonStorage()

    // TODO: Load data from files into storage.
    // For now, I start with an empty storage.

    // Setup API handlers
    apiHandlers := &api.Handlers{
        Storage: personStorage,
    }

    // Setup router
    r := chi.NewRouter()

    // A good base middleware stack
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger) // Logs the request path, method, and duration
    r.Use(middleware.Recoverer) // Recovers from panics without crashing server

    // Setup CORS
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"http://localhost:5173"}, // Your frontend's address
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300, // Maximum value not ignored by any major browsers
    }))

    // Define routes
    r.Get("/persons", apiHandlers.GetPersons)

    log.Println("Server is running on port :8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}