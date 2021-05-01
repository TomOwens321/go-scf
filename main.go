package main

import (
	"fmt"
	database "tomo/go-scf/pkg/db"
)

func main() {
	db := database.GetDB()
	p := database.Plant{Variety: "Greeny-I", SubSpecies: "Platteville"}

	// use FirstOrCreate to prevent record duplication
	db.Where(&database.Genus{Name: "Greenus"}).FirstOrCreate(&p.Genus)
	db.Where(&database.Species{Name: "plantus"}).FirstOrCreate(&p.Species)

	db.Save(&p)

	fmt.Println("Saved plant:", p.Name)
}
