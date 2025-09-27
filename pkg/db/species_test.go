package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"
)

func Test_AllSpecies(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	SeedTestData(db)

	species, err := database.AllSpecies(db)
	if err != nil {
		t.Fatalf("Failed to retrieve species: %v", err)
	}

	if len(species) == 0 {
		t.Fatal("Expected at least one species")
	}
}

func Test_GetSpecies(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	SeedTestData(db)

	species, err := database.GetSpecies(db, 1)
	if err != nil {
		t.Fatalf("Failed to retrieve species: %v", err)
	}

	if species.ID == 0 {
		t.Fatal("Expected a valid species")
	}
}

