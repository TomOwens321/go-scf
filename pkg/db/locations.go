package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Location struct {
	gorm.Model
	Name        string `gorm:"unique;index"`
	Description string
	Plants      []Plant `gorm:"many2many:plant_locations;"`
	Lattitude   float64
	Longitute   float64
	Altitude    int
	State       string
	County      string
	City        string
}

func AllLocations(db *gorm.DB) ([]Location, error) {
	var locations []Location
	err := db.Preload(clause.Associations).Find(&locations)
	return locations, err.Error
}

func GetLocation(db *gorm.DB, id uint) (Location, error) {
	var location Location
	err := db.Preload(clause.Associations).Find(&location, id)
	return location, err.Error
}

func GetLocationByName(db *gorm.DB, name string) (Location, error) {
	var location Location
	err := db.Preload(clause.Associations).Where("name = ?", name).First(&location)
	return location, err.Error
}
