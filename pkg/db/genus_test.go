package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"
)

func Test_AllGenus(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	SeedTestData(db)

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
	SeedTestData(db)

	genus, err := database.GetGenus(db, 1)
	if err != nil {
		t.Fatalf("Failed to retrieve genus: %v", err)
	}

	if genus.ID == 0 {
		t.Fatal("Expected a valid genus")
	}
}
