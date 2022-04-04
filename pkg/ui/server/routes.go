package scfui

import (
	handlers "tomo/go-scf/pkg/ui/handlers"
)

func initializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/plants", handlers.PlantsIndex)
	router.GET("/plants/:id", handlers.PlantShow)
}
