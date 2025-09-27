package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Family struct {
	gorm.Model
	Name        string `gorm:"unique;index"`
	Description string
	Plants      []Plant
}

func AllFamilies(db *gorm.DB) ([]Family, error) {
	var families []Family
	err := db.Preload(clause.Associations).Find(&families)
	return families, err.Error
}

func GetFamily(db *gorm.DB, id uint) (Family, error) {
	var family Family
	err := db.Preload(clause.Associations).First(&family, id)
	return family, err.Error
}
