package parser

import (
	"contact-hub/backend/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// LoadPersons scans a directory for .json files, parses them concurrently,
// and returns a slice of valid Person objects.
func LoadPersons(folderPath string) ([]model.Person, error) {
    files, err := os.ReadDir(folderPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read directory %s: %w", folderPath, err)
    }

    var persons []model.Person
    var wg sync.WaitGroup
    mu := &sync.Mutex{}

    for _, file := range files {
        if file.IsDir() || !strings.HasSuffix(strings.ToLower(file.Name()), ".json") {
            continue
        }

        wg.Add(1)
        go func(fileName string) {
            defer wg.Done()

            filePath := filepath.Join(folderPath, fileName)
            person, err := parseFile(filePath)
            if err != nil {
                log.Printf("WARN: Skipping file '%s': %v", fileName, err)
                return
            }

            mu.Lock()
            persons = append(persons, *person)
            mu.Unlock()
        }(file.Name())
    }

    wg.Wait()
    log.Printf("Successfully loaded %d person records.", len(persons))
    return persons, nil
}

// parseFile opens, reads, and parses a single JSON file into a Person struct.
func parseFile(filePath string) (*model.Person, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to open: %w", err)
    }
    defer file.Close()

    data, err := io.ReadAll(file)
    if err != nil {
        return nil, fmt.Errorf("failed to read: %w", err)
    }

    var person model.Person
    if err := json.Unmarshal(data, &person); err != nil {
        return nil, fmt.Errorf("malformed JSON: %w", err)
    }

    // Basic validation for required fields
    if person.FirstName == "" || person.LastName == "" {
        return nil, fmt.Errorf("missing required field(s) (firstName, lastName)")
    }
    if person.Birthday.IsZero() {
        return nil, fmt.Errorf("birthday field is missing or invalid")
    }

    return &person, nil
}