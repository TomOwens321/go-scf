package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"

	"gorm.io/gorm"
)

func seedTestLocationsData(db *gorm.DB) {
	location := database.Location{Name: "Garden", Description: "Home garden"}
	db.Create(&location)
}

func Test_AllLocations(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestLocationsData(db)

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
	seedTestLocationsData(db)

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
	seedTestLocationsData(db)

	location, err := database.GetLocationByName(db, "Garden")
	if err != nil {
		t.Fatalf("Failed to retrieve location by name: %v", err)
	}

	if location.Name != "Garden" {
		t.Fatalf("Expected location name to be 'Garden', got '%s'", location.Name)
	}
}

