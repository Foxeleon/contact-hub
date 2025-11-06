package storage

import (
	"contact-hub/backend/internal/model"
	"reflect"
	"testing"
	"time"
)

func parseTime(t *testing.T, value string) time.Time {
	ts, err := time.Parse("2006-01-02", value)
	if err != nil {
		t.Fatalf("Failed to parse time: %v", err)
	}
	return ts
}

func timePtr(t time.Time) *time.Time {
	return &t
}

func TestPersonStorage_Query(t *testing.T) {
	mockPersons := []model.Person{
		{FirstName: "John", LastName: "Doe", Birthday: parseTime(t, "1990-05-15")},
		{FirstName: "Jane", LastName: "Smith", Birthday: parseTime(t, "1985-08-22")},
		{FirstName: "Peter", LastName: "Jones", Birthday: parseTime(t, "1992-11-30")},
		{FirstName: "John", LastName: "Wick", Birthday: parseTime(t, "1964-09-02")},
	}

	storage := NewPersonStorage(mockPersons)

	testCases := []struct {
		name          string
		params        SearchParams
		expectedIDs   []string // Using a simple identifier for comparison
		expectedTotal int
	}{
		{
			name:          "No filters, first page",
			params:        SearchParams{Page: 1, PageSize: 2},
			expectedIDs:   []string{"John Doe", "Jane Smith"},
			expectedTotal: 4,
		},
		{
			name:          "No filters, second page",
			params:        SearchParams{Page: 2, PageSize: 2},
			expectedIDs:   []string{"Peter Jones", "John Wick"},
			expectedTotal: 4,
		},
		{
			name:          "Full-text search (case-insensitive)",
			params:        SearchParams{Query: "john", Page: 1, PageSize: 10},
			expectedIDs:   []string{"John Doe", "John Wick"},
			expectedTotal: 2,
		},
		{
			name:          "Date range filter (inclusive)",
			params:        SearchParams{BirthdayFrom: timePtr(parseTime(t, "1980-01-01")), BirthdayTo: timePtr(parseTime(t, "1991-01-01")), Page: 1, PageSize: 10},
			expectedIDs:   []string{"John Doe", "Jane Smith"},
			expectedTotal: 2,
		},
		{
			name:          "Combined search and date filter",
			params:        SearchParams{Query: "wick", BirthdayFrom: timePtr(parseTime(t, "1960-01-01")), BirthdayTo: timePtr(parseTime(t, "1970-01-01")), Page: 1, PageSize: 10},
			expectedIDs:   []string{"John Wick"},
			expectedTotal: 1,
		},
		{
			name:          "No results found",
			params:        SearchParams{Query: "NonExistentName", Page: 1, PageSize: 10},
			expectedIDs:   []string{},
			expectedTotal: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := storage.Query(tc.params)

			if result.Total != tc.expectedTotal {
				t.Errorf("Expected total %d, but got %d", tc.expectedTotal, result.Total)
			}

			resultIDs := make([]string, len(result.Data))
			for i, p := range result.Data {
				resultIDs[i] = p.FirstName + " " + p.LastName
			}

			if !reflect.DeepEqual(resultIDs, tc.expectedIDs) {
				t.Errorf("Expected person IDs %v, but got %v", tc.expectedIDs, resultIDs)
			}
		})
	}
}
