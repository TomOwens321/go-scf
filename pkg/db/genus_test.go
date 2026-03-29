package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"

	"gorm.io/gorm"
)

func seedTestGenusData(db *gorm.DB) {
	genus := database.Genus{Name: "Rosa", Description: "Roses"}
	db.Create(&genus)
}

func Test_AllGenus(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestGenusData(db)

	genus, err := database.AllGenus(db)
	if err != nil {
		t.Fatalf("Failed to retrieve genus: %v", err)
	}

	if len(genus) == 0 {
		t.Fatal("Expected at least one genus")
	}
}

func Test_GetGenus(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestGenusData(db)

	genus, err := database.GetGenus(db, 1)
	if err != nil {
		t.Fatalf("Failed to retrieve genus: %v", err)
	}

	if genus.ID == 0 {
		t.Fatal("Expected a valid genus")
	}
}

func Test_GetGenus_NotFound(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	_, err := database.GetGenus(db, 999) // Assuming 999 does not exist
	if err == nil {
		t.Fatal("Expected error for non-existent genus")
	}
}

func Test_GetGenusByName(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestGenusData(db)

	genus, err := database.GetGenusByName(db, "Rosa")
	if err != nil {
		t.Fatalf("Failed to retrieve genus by name: %v", err)
	}

	if genus.ID == 0 {
		t.Fatal("Expected a valid genus")
	}
}

func Test_GetGenusPlants(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestGenusData(db)
	plants := []database.Plant{
		{GenusName: "Rosa", SpeciesName: "rose"},
		{GenusName: "Rosa", SpeciesName: "wildrose"},
	}
	db.Create(&plants)

	genus, err := database.GetGenus(db, 1)
	if err != nil {
		t.Fatalf("Failed to retrieve genus: %v", err)
	}

	if len(genus.Plants) != len(plants) {
		t.Fatalf("Expected %v plants in genus, got %v", len(plants), len(genus.Plants))
	}
}
