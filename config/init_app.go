package config

import (
	"log"

	"github.com/LucasFreitasRocha/probe-challenge-go/config/database"
	"github.com/LucasFreitasRocha/probe-challenge-go/config/logger"
	"github.com/LucasFreitasRocha/probe-challenge-go/controller"
	"github.com/LucasFreitasRocha/probe-challenge-go/repository"
	"github.com/LucasFreitasRocha/probe-challenge-go/routes"
	"github.com/LucasFreitasRocha/probe-challenge-go/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitApp() {
	godotenv.Load()
	logger.Info("database connection started")
	db, err := database.Connect()
	if err != nil {
		logger.Error("Failed to connect to the database", err)
	}

	router := gin.Default()
	probeService := service.NewProbeService(repository.NewProbeRepository(db))
	routes.SetupProbeRoutes(router, initProbeController(probeService))
	routes.SetupCommandRoutes(router, initCommandController(probeService))
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	} 

}

func initCommandController(probeService service.ProbeService) controller.CommandController {
	commandService := service.NewCommandService(probeService)
	return controller.NewCommandController(commandService)
}


func initProbeController(probeService service.ProbeService) controller.ProbeController {

	return controller.NewProbeController(probeService)
}
