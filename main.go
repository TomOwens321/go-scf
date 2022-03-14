package main

import (
	"fmt"
	database "tomo/go-scf/pkg/db"
)

func main() {
	db := database.GetDB()

	p := database.Plant{GenusName: "Greenus", SpeciesName: "plantus", SubSpecies: "sub_green", FamilyName: "Greens"}
	pName := p.FullName()

	// use FirstOrCreate to prevent record duplication
	db.Where(&database.Plant{Name: pName}).FirstOrCreate(&p)

	fmt.Println("Plant name: ", p.Name)

	p.CommonName = "Green Plant"
	db.Save(&p)

	fmt.Println("Saved plant:", p.Name)
}
