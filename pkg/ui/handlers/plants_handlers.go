package scfui

import (
	"net/http"
	db "tomo/go-scf/pkg/db"

	"github.com/gin-gonic/gin"
)

func PlantsIndex(c *gin.Context) {
	plants := db.AllPlants()
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"plants.html",
		// Pass the data that the page uses
		gin.H{
			"title":  "Plant List",
			"plants": plants,
		},
	)

}
