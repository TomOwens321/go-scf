package scfui

import (
	handlers "tomo/go-scf/pkg/ui/handlers"
)

func initializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/plants", handlers.PlantsIndex)
	router.GET("/plants/:id", handlers.PlantShow)
	router.GET("/plants/new", handlers.PlantNew)
	router.GET("/plants/:id/edit", handlers.PlantEdit)
	router.POST("plants/create", handlers.PlantCreate)
}
