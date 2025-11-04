package main

import (
	"fmt"
	"log"
	"contact-hub/backend/internal/parser"
)

func main() {
    fmt.Println("Launching the service Contact Hub...")

    // Example of calling the parser stub
    persons, err := parser.LoadPersons("./data")
    if err != nil {
    	log.Fatalf("Error when uploading data: %v", err)
    }

    log.Printf("The parser stub has been successfully invoked. Uploaded %d records.", len(persons))
    // TODO: Add storage initialization and HTTP server startup.
}