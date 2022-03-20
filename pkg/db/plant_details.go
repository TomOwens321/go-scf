package database

import (
	"gorm.io/gorm"
)

type PlantDetail struct {
	gorm.Model
	PlantID     uint
	Description string
	Altitude    string
	Zone        uint
}

func init() {
	db := GetDB()
	db.AutoMigrate(&PlantDetail{})
}
