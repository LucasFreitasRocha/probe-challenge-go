package main

import (
	"log"

	"github.com/LucasFreitasRocha/probe-challenge-go/config/database"
	"github.com/LucasFreitasRocha/probe-challenge-go/config/logger"
	"github.com/LucasFreitasRocha/probe-challenge-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/LucasFreitasRocha/probe-challenge-go/controller"
	"github.com/LucasFreitasRocha/probe-challenge-go/repository"
	"github.com/LucasFreitasRocha/probe-challenge-go/service"
	"gorm.io/gorm"
)

func main() {
	logger.Info("About to start user application")

	godotenv.Load()
	
	logger.Info("database connection started")
	db, err := database.Connect()
	if err != nil {
		logger.Error("Failed to connect to the database", err)
	}

	probeController := initDependencies(db)

	// Create a new gin router
	router := gin.Default()
	routes.SetupProbeRoutes(router, probeController)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	} // Listen and serve on port 8080
}


func initDependencies(db *gorm.DB) controller.ProbeController {
	repo := repository.NewProbeRepository(db)
	service := service.NewProbeService(repo)
	return controller.NewProbeController(service)
}