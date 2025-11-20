package parser

import (
	"bytes"
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

			parsedPersons, err := parseFile(filePath)
			if err != nil {
				log.Printf("WARN: Skipping file '%s': %v", fileName, err)
				return
			}

			if len(parsedPersons) > 0 {
				mu.Lock()
				persons = append(persons, parsedPersons...)
				mu.Unlock()
			}
		}(file.Name())
	}

	wg.Wait()
	log.Printf("Successfully loaded %d person records.", len(persons))
	return persons, nil
}

func validatePerson(person model.Person) error {
	if person.FirstName == "" || person.LastName == "" {
		return fmt.Errorf("missing required field(s) (firstName, lastName)")
	}
	if person.Birthday.IsZero() {
		return fmt.Errorf("birthday field is missing or invalid")
	}
	return nil
}

func parseFile(filePath string) ([]model.Person, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read: %w", err)
	}

	trimmedData := bytes.TrimSpace(data)
	if len(trimmedData) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	var validPersons []model.Person

	if trimmedData[0] == '[' {
		var persons []model.Person
		if err := json.Unmarshal(data, &persons); err != nil {
			return nil, fmt.Errorf("malformed JSON array: %w", err)
		}

		for _, p := range persons {
			if err := validatePerson(p); err == nil {
				validPersons = append(validPersons, p)
			} else {
				log.Printf("WARN: Skipping invalid record in file '%s': %v", filepath.Base(filePath), err)
			}
		}
	} else if trimmedData[0] == '{' {
		var person model.Person
		if err := json.Unmarshal(data, &person); err != nil {
			return nil, fmt.Errorf("malformed JSON object: %w", err)
		}

		if err := validatePerson(person); err == nil {
			validPersons = append(validPersons, person)
		} else {
			return nil, fmt.Errorf("invalid person record: %w", err)
		}
	} else {
		return nil, fmt.Errorf("invalid JSON format: must be an object or an array")
	}

	return validPersons, nil
}
