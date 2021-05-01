package database

import (
	"gorm.io/gorm"
)

type Genus struct {
	gorm.Model
	Name        string `gorm:"unique"`
	Description string
}

func init() {
	db := GetDB()
	db.AutoMigrate(&Genus{})
}
