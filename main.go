package main

import (
	"fmt"
	database "tomo/go-scf/pkg/db"
)

func main() {
	db := database.GetDB()

	p := database.Plant{GenusName: "Greenus", SpeciesName: "plantus"}
	pName := p.FullName()

	// use FirstOrCreate to prevent record duplication
	db.Where(&database.Plant{Name: pName}).Preload("PlantDetail").FirstOrCreate(&p)

	fmt.Println("Plant name: ", p.Name)

	p.CommonName = "Green Plant"
	p.PlantDetail.Description = "This is a lovely green plant."
	db.Save(&p)

	fmt.Println("Saved plant:", p.Name)

	g, err := database.GetGenus(p.Genus.ID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(g.Name, g.Plants[0].Name)
	}

	s, err := database.GetSpecies(p.Species.ID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s.Name, s.Plants[0].Name)
	}
}
