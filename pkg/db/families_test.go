package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"
)

func Test_AllFamilies(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	// Add test data
	SeedTestData(db)

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
	SeedTestData(db)

	family, err := database.GetFamily(db, 1)
	if err != nil {
		t.Fatalf("Failed to retrieve family: %v", err)
	}

	if family.ID == 0 {
		t.Fatal("Expected a valid family")
	}
}
