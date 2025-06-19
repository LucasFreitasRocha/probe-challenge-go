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
	routes.SetupProbeRoutes(router, initControllers(db))
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	} 

}


func initControllers(db *gorm.DB) controller.ProbeController {
	probeRepository := repository.NewProbeRepository(db)
	probeService := service.NewProbeService(probeRepository)
	return controller.NewProbeController(probeService)
}
