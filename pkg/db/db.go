package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbPath = "scf.db"

func connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{FullSaveAssociations: true})
	if err != nil {
		panic("Failed to open the database")
	}
	// Migrate the schema
	if err := db.AutoMigrate(
		&Family{},
		&Genus{},
		&Species{},
		&Location{},
		&PlantDetail{},
		&Plant{}); err != nil {
		panic("Failed to migrate the database")
	}

	return db
}

func GetDB() *gorm.DB {
	return connect()
}
