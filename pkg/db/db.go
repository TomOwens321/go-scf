package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbPath = "db/scf.db"

var db *gorm.DB

func init() {
	db = connect()
}

func connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("Failed to open the database")
	}

	return db
}

func GetDB() *gorm.DB {
	return db
}
