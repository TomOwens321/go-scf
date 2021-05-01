package database

import "gorm.io/gorm"

type Species struct {
	gorm.Model
	Name        string	`gorm:"unique"`
	Description string
}

func init() {
	db := GetDB()
	db.AutoMigrate(&Species{})
}
