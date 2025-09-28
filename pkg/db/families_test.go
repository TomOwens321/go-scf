package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"

	"gorm.io/gorm"
)

func seedTestFamiliesData(db *gorm.DB) {
	family := database.Family{Name: "Rosaceae", Description: "Rose family"}
	db.Create(&family)
}

func Test_AllFamilies(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestFamiliesData(db)

	families, err := database.AllFamilies(db)
	if err != nil {
		t.Fatalf("Failed to retrieve families: %v", err)
	}

	if len(families) == 0 {
		t.Fatal("Expected at least one family")
	}
}

func Test_GetFamily(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestFamiliesData(db)

	family, err := database.GetFamily(db, 1)
	if err != nil {
		t.Fatalf("Failed to retrieve family: %v", err)
	}

	if family.ID == 0 {
		t.Fatal("Expected a valid family")
	}
}

func Test_GetFamilyPlants(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestFamiliesData(db)
	plants := []database.Plant{
		{GenusName: "Rosa", SpeciesName: "Rose", FamilyID: 1},
		{GenusName: "Malus", SpeciesName: "Apple", FamilyID: 1},
	}
	db.Create(&plants)

	family, err := database.GetFamily(db, 1)
	if err != nil {
		t.Fatalf("Failed to retrieve family: %v", err)
	}

	if len(family.Plants) == 0 {
		t.Fatal("Expected at least one plant in the family")
	}

	found := false
	for _, plant := range family.Plants {
		if plant.Name == "Rosa Rose" {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Expected to find plant 'Rosa Rose' in the family")
	}

	for _, plant := range family.Plants {
		if plant.FamilyID != int(family.ID) {
			t.Fatalf("Expected plant FamilyID to be %d, got %d", family.ID, plant.FamilyID)
		}
	}
}

func Test_GetFamilyNotFound(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	_, err := database.GetFamily(db, 999) // Assuming 999 does not exist
	if err == nil {
		t.Fatal("Expected error for non-existent family")
	}
}
