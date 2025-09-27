package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"

	"gorm.io/gorm"
)

func seedTestSpeciesData(db *gorm.DB) {
	species := database.Species{Name: "rubiginosa", Description: "Sweet briar rose"}
	db.Create(&species)
}
func Test_AllSpecies(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestSpeciesData(db)

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
	seedTestSpeciesData(db)

	species, err := database.GetSpecies(db, 1)
	if err != nil {
		t.Fatalf("Failed to retrieve species: %v", err)
	}

	if species.ID == 0 {
		t.Fatal("Expected a valid species")
	}

	if species.Name != "rubiginosa" {
		t.Fatalf("Expected species name 'rubiginosa', got '%s'", species.Name)
	}
}
