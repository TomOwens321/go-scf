package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Plant struct {
	gorm.Model
	Name        string `gorm:"unique;index"`
	CommonName  string
	Genus       Genus
	Species     Species
	Family      Family
	PlantDetail PlantDetail
	SubSpecies  string
	Variety     string
	GenusID     int
	SpeciesID   int
	FamilyID    int
	GenusName   string `gorm:"-"`
	SpeciesName string `gorm:"-"`
	FamilyName  string `gorm:"-"`
}

type PlantName struct {
	Genus      string
	Species    string
	SubSpecies string
	Variety    string
}

func init() {
	db := GetDB()
	db.AutoMigrate(&Plant{})
}

func (p *Plant) FullName() string {
	rslt := p.GenusName + " " + p.SpeciesName

	if len(p.SubSpecies) > 0 {
		rslt += " ssp: " + p.SubSpecies
	}

	if len(p.Variety) > 0 {
		rslt += " var: " + p.Variety
	}

	return rslt
}

func AllPlants() []Plant {
	var plants []Plant
	db.Preload(clause.Associations).Find(&plants)
	return plants
}

func (p *Plant) BeforeSave(tx *gorm.DB) (err error) {
	p.Name = p.FullName()
	db.Where(Genus{Name: p.GenusName}).FirstOrCreate(&p.Genus)
	db.Where(Species{Name: p.SpeciesName}).FirstOrCreate(&p.Species)
	if len(p.FamilyName) > 0 {
		db.Where(Family{Name: p.FamilyName}).FirstOrCreate(&p.Family)
	}
	return
}
