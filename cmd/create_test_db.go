package main

import (
	"fmt"
	database "tomo/go-scf/pkg/db"
)

func main() {

	loc, _ := createTestLocation()
	p, _ := createTestPlant(&loc)

	fmt.Println("Plant name: ", p.Name)
	fmt.Println("Plant common name: ", p.CommonName)
	fmt.Println("Plant Family: ", p.FamilyName)
	fmt.Println("Plant Genus: ", p.GenusName)
	fmt.Println("Plant Species: ", p.SpeciesName)
	fmt.Println("Plant description: ", p.PlantDetail.Description)
	fmt.Println("Plant location 1: ", p.Locations[0].Name)
}

func createTestPlant(location *database.Location) (database.Plant, error) {
	db := database.GetDB()

	p := database.Plant{GenusName: "Greenus", SpeciesName: "plantus", FamilyName: "The Greens", Variety: "Platteville", SubSpecies: "greeny"}
	pName := p.FullName()

	// use FirstOrCreate to prevent record duplication
	db.Where(&database.Plant{Name: pName}).Preload("PlantDetail").FirstOrCreate(&p)

	// fmt.Println("Plant name: ", p.Name)

	p.CommonName = "Green Plant"
	p.PlantDetail.Description = "This is a lovely green plant."
	p.PlantDetail.FlowerColor = "Green"
	p.PlantDetail.FlowerTime = "Spring"
	p.PlantDetail.Hardiness = "Hardy"
	p.PlantDetail.Height = 1.5 // meters
	p.PlantDetail.Spread = 0.5 // meters
	p.PlantDetail.Zone = 5
	p.PlantDetail.Altitude = 1000 // meters
	p.PlantDetail.Climate = "Temperate"
	p.Locations = append(p.Locations, *location)
	db.Save(&p)

	// fmt.Println("Saved plant:", p.Name)
	return p, nil
}

func createTestLocation() (database.Location, error) {
	db := database.GetDB()

	loc := database.Location{Name: "Test Location", Description: "This is a test location."}
	db.Where(&database.Location{Name: loc.Name}).Preload("Plants").FirstOrCreate(&loc)

	// fmt.Println("Location name: ", loc.Name)

	loc.Description = "This is a test location."
	db.Save(&loc)

	// fmt.Println("Saved location:", loc.Name)
	return loc, nil
}
