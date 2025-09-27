package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"
)

func Test_AllPlants(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	SeedTestData(db)

	plants := database.AllPlants(db)
	if len(plants) == 0 {
		t.Fatal("Expected at least one plant")
	}
}

func Test_GetPlantByID(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	SeedTestData(db)

	plant := database.GetPlantByID(db, 1)
	if plant.ID == 0 {
		t.Fatal("Expected a valid plant")
	}

	if plant.Genus.Name != "Rosa" {
		t.Fatalf("Expected genus name to be 'Rosa', got '%s'", plant.Genus.Name)
	}
	if plant.Family.Name != "Rosaceae" {
		t.Fatalf("Expected family name to be 'Rosaceae', got '%s'", plant.Family.Name)
	}
	if len(plant.Locations) == 0 || plant.Locations[0].Name != "Test Location" {
		t.Fatalf("Expected location name to be 'Test Location', got '%v'", plant.Locations)
	}
}
