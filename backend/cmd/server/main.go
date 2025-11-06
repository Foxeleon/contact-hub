package main

import (
	"contact-hub/backend/internal/api"
	"contact-hub/backend/internal/parser"
	"contact-hub/backend/internal/storage"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	log.Println("Starting Contact Hub service...")

	// 1. Load person data from files using the parser
	loadedPersons, err := parser.LoadPersons("./data")
	if err != nil {
		// Log a fatal error because the service is useless without data
		log.Fatalf("FATAL: Failed to load initial data: %v", err)
	}

	// 2. Create storage instance with the loaded data
	personStorage := storage.NewPersonStorage(loadedPersons)

	// Setup API handlers
	apiHandlers := &api.Handlers{
		Storage: personStorage,
	}

	// Setup router
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Enable CORS only if it is running in development mode.
	if os.Getenv("APP_ENV") == "development" {
		log.Println("Development mode: Enabling CORS for localhost:5173")
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"http://localhost:5173"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any browser
		}))
	}

	// Group all API routes under the /api prefix
	r.Route("/api", func(r chi.Router) {
		// This now correctly maps to GET /api/persons
		r.Get("/persons", apiHandlers.GetPersons)

		// to add more routes in the future, they go here:
		// r.Get("/stats", apiHandlers.GetStats)
	})

	log.Println("Server is running on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
