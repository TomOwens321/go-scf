package database

import (
	"gorm.io/gorm"
)

type Plant struct {
	gorm.Model
	Name        string `gorm:"unique;index"`
	CommonNames string
	Genus       Genus
	Species     Species
	SubSpecies  string
	Variety     string
	GenusID     int
	SpeciesID   int
}

func init() {
	db := GetDB()
	db.AutoMigrate(&Plant{})
}

func (p *Plant) setFullName() string {
	rslt := p.Genus.Name + " " + p.Species.Name
	if len(p.SubSpecies) > 0 {
		rslt += " ssp: " + p.SubSpecies
	}
	if len(p.Variety) > 0 {
		rslt += " var: " + p.Variety
	}

	return rslt
}

func (p *Plant) BeforeSave(tx *gorm.DB) (err error) {
	name := p.setFullName()
	p.Name = name
	return
}
