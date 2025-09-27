package database_test

import (
	"testing"
	database "tomo/go-scf/pkg/db"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		t.Fatalf("Failed to open the test database: %v", err)
	}
	// Migrate the schema
	if err := db.AutoMigrate(
		&database.Family{},
		&database.Genus{},
		&database.Species{},
		&database.Location{},
		&database.PlantDetail{},
		&database.Plant{}); err != nil {
		t.Fatalf("Failed to migrate the test database: %v", err)
	}
	return db
}

func TeardownTestDB(t *testing.T, db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		t.Fatalf("Failed to close the test database: %v", err)
	}
}

func SeedTestData(db *gorm.DB) {
	family := database.Family{Name: "Rosaceae", Description: "Rose family"}
	db.Create(&family)

	genus := database.Genus{Name: "Rosa", Description: "Roses"}
	db.Create(&genus)

	species := database.Species{Name: "rubiginosa", Description: "Sweet briar rose"}
	db.Create(&species)

	location := database.Location{Name: "Test Location", Description: "A location for testing"}
	db.Create(&location)

	plant_detail := database.PlantDetail{
		Altitude:    100,
		Climate:     "Temperate",
		Description: "A test plant detail",
		FlowerColor: "Red",
		FlowerTime:  "Spring",
		Hardiness:   "Hardy",
		Height:      1.5,
		Spread:      0.5,
		Zone:        5,
	}
	db.Create(&plant_detail)

	plant := database.Plant{
		CommonName:  "Test Rose",
		GenusName:   genus.Name,
		FamilyName: family.Name,
		Locations:   []database.Location{location},
	}
	db.Create(&plant)
}

func Test_GetDB(t *testing.T) {
	db := SetupTestDB(t)
	defer TeardownTestDB(t, db)

	if db == nil {
		t.Fatal("Expected a valid database connection")
	}
}
