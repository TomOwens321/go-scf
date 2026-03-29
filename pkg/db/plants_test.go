package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"

	"gorm.io/gorm"
)

func seedTestPlantsData(db *gorm.DB) {
	plant := database.Plant{
		CommonName: "Test Rose",
		GenusName:  "genus1",
		FamilyName: "family1",
		SpeciesName: "species1",
		Locations:  []database.Location{{Name: "Test Location", Description: "A location for testing"}},
	}
	db.Create(&plant)
}

func Test_AllPlants(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestPlantsData(db)

	plants := database.AllPlants(db)
	if len(plants) == 0 {
		t.Fatal("Expected at least one plant")
	}
}

func Test_GetPlantByID(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestPlantsData(db)

	plant := database.GetPlantByID(db, 1)
	if plant.ID == 0 {
		t.Fatal("Expected a valid plant")
	}

	if plant.Genus.Name != "genus1" {
		t.Fatalf("Expected genus name to be 'genus1', got '%s'", plant.Genus.Name)
	}
	if plant.Family.Name != "family1" {
		t.Fatalf("Expected family name to be 'family1', got '%s'", plant.Family.Name)
	}
	if len(plant.Locations) == 0 || plant.Locations[0].Name != "Test Location" {
		t.Fatalf("Expected location name to be 'Test Location', got '%v'", plant.Locations)
	}
}

func Test_GetPlantByName(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	seedTestPlantsData(db)

	plant, err := database.GetPlantByName(db, "genus1 species1")
	if err != nil {
		t.Fatalf("Failed to retrieve plant by name: %v", err)
	}

	if plant.ID == 0 {
		t.Fatal("Expected a valid plant")
	}
}

func Test_PlantNames(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	expect := []string{
		"Greenus plantus",
		"Greenus plantus ssp: subplantus",
		"Greenus plantus var: varietus",
		"Greenus plantus ssp: subplantus var: varietus",
	}

	// Add test data
	plants := []database.Plant{
		{
			GenusName:   "Greenus",
			SpeciesName: "plantus",
		},
		{
			GenusName:   "Greenus",
			SpeciesName: "plantus",
			SubSpecies:  "subplantus",
		},
		{
			GenusName:   "Greenus",
			SpeciesName: "plantus",
			Variety:     "varietus",
		},
		{
			GenusName:   "Greenus",
			SpeciesName: "plantus",
			SubSpecies:  "subplantus",
			Variety:     "varietus",
		},
	}
	db.Create(&plants)

	got := database.AllPlants(db)
	if len(got) != 4 {
		t.Fatal("Expected 4 plants")
	}
	for i, plant := range got {
		if plant.Name != expect[i] {
			t.Fatalf("Expected full name to be '%s', got '%s'", expect[i], plant.Name)
		}
	}
}

