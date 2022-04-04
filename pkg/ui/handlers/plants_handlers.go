package scfui

import (
	"net/http"
	"strconv"
	db "tomo/go-scf/pkg/db"

	"github.com/gin-gonic/gin"
)

func PlantsIndex(c *gin.Context) {
	plants := db.AllPlants()
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"plants/index.html",
		// Pass the data that the page uses
		gin.H{
			"title":  "Plant List",
			"plants": plants,
		},
	)

}

func PlantShow(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	plant := db.GetPlantByID(uint(id))
	c.HTML(
		http.StatusOK,
		"plants/show.html",
		gin.H{
			"title": plant.Name,
			"plant": plant,
		},
	)
}
