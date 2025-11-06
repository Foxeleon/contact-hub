package storage

import (
	"contact-hub/backend/internal/model"
	"reflect"
	"testing"
	"time"
)

// parseTime is a test helper function to parse time strings, failing the test on error.
func parseTime(t *testing.T, value string) time.Time {
	ts, err := time.Parse("2006-01-02", value)
	if err != nil {
		t.Fatalf("Failed to parse time: %v", err)
	}
	return ts
}

// timePtr is a test helper function to get a pointer to a time.Time value.
func timePtr(t time.Time) *time.Time {
	return &t
}

func TestPersonStorage_Query(t *testing.T) {
	// A consistent set of mock data for all test cases.
	mockPersons := []model.Person{
		{FirstName: "John", LastName: "Doe", Birthday: parseTime(t, "1990-05-15")},
		{FirstName: "Jane", LastName: "Smith", Birthday: parseTime(t, "1985-08-22")},
		{FirstName: "Peter", LastName: "Jones", Birthday: parseTime(t, "1992-11-30")},
		{FirstName: "John", LastName: "Wick", Birthday: parseTime(t, "1964-09-02")},
		{FirstName: "Sarah", LastName: "Connor", Birthday: parseTime(t, "1965-05-13")},
		{FirstName: "Martin", LastName: "McFly", Birthday: parseTime(t, "1968-06-12")},
	}

	storage := NewPersonStorage(mockPersons)

	// A comprehensive suite of test cases covering happy paths and edge cases.
	testCases := []struct {
		name          string
		params        SearchParams
		expectedIDs   []string
		expectedTotal int
	}{
		// --- Standard Cases ---
		{
			// Sorted data: Jane Smith, John Doe, John Wick, Martin McFly, Peter Jones, Sarah Connor
			name:          "Should return the first page when no filters are applied",
			params:        SearchParams{Page: 1, PageSize: 3},
			expectedIDs:   []string{"Jane Smith", "John Doe", "John Wick"},
			expectedTotal: 6,
		},
		{
			name:          "Should return the second page when no filters are applied",
			params:        SearchParams{Page: 2, PageSize: 3},
			expectedIDs:   []string{"Martin McFly", "Peter Jones", "Sarah Connor"},
			expectedTotal: 6,
		},
		{
			// Filtered: [John Doe, John Wick]. Sorted: no change.
			name:          "Should filter by first name (case-insensitive)",
			params:        SearchParams{Query: "john", Page: 1, PageSize: 10},
			expectedIDs:   []string{"John Doe", "John Wick"},
			expectedTotal: 2,
		},
		{
			// Filtered: [Sarah Connor].
			name:          "Should filter by last name",
			params:        SearchParams{Query: "connor", Page: 1, PageSize: 10},
			expectedIDs:   []string{"Sarah Connor"},
			expectedTotal: 1,
		},
		{
			// Filtered: [John Doe (1990), Peter Jones (1992)]. Sorted: John, Peter.
			name:          "Should filter by date range (from only)",
			params:        SearchParams{BirthdayFrom: timePtr(parseTime(t, "1990-01-01")), Page: 1, PageSize: 10},
			expectedIDs:   []string{"John Doe", "Peter Jones"},
			expectedTotal: 2,
		},
		{
			// Filtered: [John Wick (1964), Sarah Connor (1965)]. Sorted: John, Sarah.
			name:          "Should filter by date range (to only)",
			params:        SearchParams{BirthdayTo: timePtr(parseTime(t, "1965-12-31")), Page: 1, PageSize: 10},
			expectedIDs:   []string{"John Wick", "Sarah Connor"},
			expectedTotal: 2,
		},
		{
			// Filtered: [John Doe (1990), Jane Smith (1985)]. Sorted: Jane, John.
			name:          "Should handle combined text and date filters",
			params:        SearchParams{Query: "j", BirthdayFrom: timePtr(parseTime(t, "1980-01-01")), BirthdayTo: timePtr(parseTime(t, "1991-01-01")), Page: 1, PageSize: 10},
			expectedIDs:   []string{"Jane Smith", "John Doe"},
			expectedTotal: 2,
		},
		{
			// Filtered: [Wick(64), Connor(65), McFly(68), Smith(85)].
			// Sorted: Jane Smith, John Wick, Martin McFly, Sarah Connor.
			// Page 2: [Martin McFly, Sarah Connor].
			name:          "Should paginate correctly on filtered results",
			params:        SearchParams{BirthdayTo: timePtr(parseTime(t, "1989-12-31")), Page: 2, PageSize: 2},
			expectedIDs:   []string{"Martin McFly", "Sarah Connor"},
			expectedTotal: 4,
		},
		// --- Edge Cases ---
		{
			name:          "Should return no results for a query that matches nothing",
			params:        SearchParams{Query: "NonExistentName", Page: 1, PageSize: 10},
			expectedIDs:   []string{},
			expectedTotal: 0,
		},
		{
			name:          "Should return an empty slice for a page that is out of bounds",
			params:        SearchParams{Page: 10, PageSize: 10},
			expectedIDs:   []string{},
			expectedTotal: 6,
		},
		{
			name: "Should handle queries on an empty storage",
			// This test case uses a separate, empty storage instance.
			params:        SearchParams{Page: 1, PageSize: 10},
			expectedIDs:   []string{},
			expectedTotal: 0,
		},
		{
			name:          "Should treat an invalid page number (e.g., 0) as page 1",
			params:        SearchParams{Page: 0, PageSize: 3},
			expectedIDs:   []string{"Jane Smith", "John Doe", "John Wick"}, // Expects the first page
			expectedTotal: 6,
		},
		{
			name:          "Should return no results for an invalid date range (from > to)",
			params:        SearchParams{BirthdayFrom: timePtr(parseTime(t, "2000-01-01")), BirthdayTo: timePtr(parseTime(t, "1990-01-01")), Page: 1, PageSize: 10},
			expectedIDs:   []string{},
			expectedTotal: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Use the main storage by default
			currentStorage := storage

			// For the specific empty storage test, create a new empty instance
			if tc.name == "Should handle queries on an empty storage" {
				currentStorage = NewPersonStorage([]model.Person{})
			}

			// Execute the query
			result := currentStorage.Query(tc.params)

			// Assert the total number of records
			if result.Total != tc.expectedTotal {
				t.Errorf("Expected total %d, but got %d", tc.expectedTotal, result.Total)
			}

			// Convert result data to a comparable format (slice of strings)
			resultIDs := make([]string, len(result.Data))
			for i, p := range result.Data {
				resultIDs[i] = p.FirstName + " " + p.LastName
			}

			// Assert the content of the returned data
			if !reflect.DeepEqual(resultIDs, tc.expectedIDs) {
				t.Errorf("Expected person IDs %v, but got %v", tc.expectedIDs, resultIDs)
			}
		})
	}
}
