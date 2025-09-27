package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"
)

func Test_AllLocations(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	SeedTestData(db)

	locations, err := database.AllLocations(db)
	if err != nil {
		t.Fatalf("Failed to retrieve locations: %v", err)
	}

	if len(locations) == 0 {
		t.Fatal("Expected at least one location")
	}
}

func Test_GetLocation(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	SeedTestData(db)

	location, err := database.GetLocation(db, 1)
	if err != nil {
		t.Fatalf("Failed to retrieve location: %v", err)
	}

	if location.ID == 0 {
		t.Fatal("Expected a valid location")
	}
}

func Test_GetLocationByName(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	SeedTestData(db)

	location, err := database.GetLocationByName(db, "Test Location")
	if err != nil {
		t.Fatalf("Failed to retrieve location by name: %v", err)
	}

	if location.Name != "Test Location" {
		t.Fatalf("Expected location name to be 'Test Location', got '%s'", location.Name)
	}
}
