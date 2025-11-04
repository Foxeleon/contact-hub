package main

import (
	"contact-hub/backend/internal/parser"
	"log"
	"os"
)

func main() {
	// Set up a structured logger
	// We use the standard 'log' package, but configure it for better output.
	logger := log.New(os.Stdout, "CONTACT-HUB: ", log.LstdFlags|log.Lshortfile)

	logger.Println("Service starting...")

	// --- Data Loading Step ---
	// We call the parser to load person data from the relative 'data' directory.
	persons, err := parser.LoadPersons("./data")
	if err != nil {
		// If the directory cannot be read, it's a fatal error.
		logger.Fatalf("FATAL: Failed to load person data: %v", err)
	}

	logger.Printf("Successfully loaded %d person records.", len(persons))

	// --- Next Steps (Placeholders) ---
	// TODO: Initialize the in-memory storage with the loaded persons.
	// TODO: Initialize the Chi router and start the HTTP server.
}