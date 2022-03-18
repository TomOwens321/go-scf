package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Genus struct {
	gorm.Model
	Name        string `gorm:"unique;index"`
	Description string
	Plants      []Plant
}

func init() {
	db := GetDB()
	db.AutoMigrate(&Genus{})
}

func AllGenus() ([]Genus, error) {
	var genus []Genus
	err := db.Preload(clause.Associations).Find(&genus)
	return genus, err.Error
}

func GetGenus(id uint) (Genus, error) {
	var genus Genus
	err := db.Preload(clause.Associations).First(&genus, id)
	return genus, err.Error
}
