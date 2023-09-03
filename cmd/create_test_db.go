package main

import (
	"fmt"
	database "tomo/go-scf/pkg/db"
)

func main() {

	loc, _ := createTestLocation()
	createTestPlant(&loc)

}

func createTestPlant(location *database.Location) {
	db := database.GetDB()

	p := database.Plant{GenusName: "Greenus", SpeciesName: "plantus", FamilyName: "The Greens", Variety: "Platteville", SubSpecies: "greeny"}
	pName := p.FullName()

	// use FirstOrCreate to prevent record duplication
	db.Where(&database.Plant{Name: pName}).Preload("PlantDetail").FirstOrCreate(&p)

	fmt.Println("Plant name: ", p.Name)

	p.CommonName = "Green Plant"
	p.PlantDetail.Description = "This is a lovely green plant."
	p.Locations = append(p.Locations, *location)
	db.Save(&p)

	fmt.Println("Saved plant:", p.Name)
}

func createTestLocation() (database.Location, error) {
	db := database.GetDB()

	loc := database.Location{Name: "Test Location", Description: "This is a test location."}
	db.Where(&database.Location{Name: loc.Name}).Preload("Plants").FirstOrCreate(&loc)

	fmt.Println("Location name: ", loc.Name)

	loc.Description = "This is a test location."
	db.Save(&loc)

	fmt.Println("Saved location:", loc.Name)
	return loc, nil
}
