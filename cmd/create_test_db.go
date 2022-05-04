package main

import (
	"fmt"
	database "tomo/go-scf/pkg/db"
)

func main() {
	db := database.GetDB()

	p := database.Plant{GenusName: "Greenus", SpeciesName: "plantus", FamilyName: "The Greens", Variety: "Platteville", SubSpecies: "greeny"}
	pName := p.FullName()

	// use FirstOrCreate to prevent record duplication
	db.Where(&database.Plant{Name: pName}).Preload("PlantDetail").FirstOrCreate(&p)

	fmt.Println("Plant name: ", p.Name)

	p.CommonName = "Green Plant"
	p.PlantDetail.Description = "This is a lovely green plant."
	db.Save(&p)

	fmt.Println("Saved plant:", p.Name)

}
