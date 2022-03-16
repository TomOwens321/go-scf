package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Species struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
	Plants      []Plant
}

func init() {
	db := GetDB()
	db.AutoMigrate(&Species{})
}

func AllSpecies() ([]Species, error) {
	var species []Species
	err := db.Preload(clause.Associations).Find(&species)
	return species, err.Error
}

func GetSpecies(id uint) (Species, error) {
	var species Species
	err := db.Preload(clause.Associations).First(&species, id)
	return species, err.Error
}
