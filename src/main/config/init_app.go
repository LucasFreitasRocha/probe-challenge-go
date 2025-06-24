package config

import (
	"log"

	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config/database"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config/logger"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/controller"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/repository"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/routes"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func InitApp() {
	godotenv.Load()
	logger.Info("database connection started")
	db, err := database.Connect()
	if err != nil {
		logger.Error("Failed to connect to the database", err)
	}

	router := gin.Default()
	probeService := InitProbeService(db)
	routes.SetupProbeRoutes(router, InitProbeController(probeService))
	routes.SetupCommandRoutes(router, InitCommandController(probeService))
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	} 

}

func InitProbeService(database *gorm.DB) service.ProbeService {
	return service.NewProbeService(repository.NewProbeRepository(database))
}

func InitCommandController(probeService service.ProbeService) controller.CommandController {
	commandService := service.NewCommandService(probeService)
	return controller.NewCommandController(commandService)
}


func InitProbeController(probeService service.ProbeService) controller.ProbeController {

	return controller.NewProbeController(probeService)
}
