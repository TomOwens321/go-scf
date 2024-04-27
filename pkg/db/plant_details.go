package database

import (
	"gorm.io/gorm"
)

type PlantDetail struct {
	gorm.Model
	PlantID     uint
	Altitude    uint
	Climate     string
	Description string
	FlowerColor string
	FlowerTime  string
	Hardiness   string
	Height      float32
	Spread      float32
	Zone        uint
}

func init() {
	db := GetDB()
	db.AutoMigrate(&PlantDetail{})
}
