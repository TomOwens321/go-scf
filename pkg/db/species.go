package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Species struct {
	gorm.Model
	Name        string `gorm:"unique;index"`
	Description string
	Plants      []Plant
}

func AllSpecies(db *gorm.DB) ([]Species, error) {
	var species []Species
	err := db.Preload(clause.Associations).Find(&species)
	return species, err.Error
}

func GetSpecies(db *gorm.DB, id uint) (Species, error) {
	var species Species
	err := db.Preload(clause.Associations).First(&species, id)
	return species, err.Error
}
