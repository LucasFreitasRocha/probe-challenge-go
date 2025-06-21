package routes

import (
	"github.com/LucasFreitasRocha/probe-challenge-go/controller"
	"github.com/gin-gonic/gin"
)

func SetupProbeRoutes(router *gin.Engine, controller controller.ProbeController) {

	probeGroup := router.Group("/probes")
	{
		probeGroup.POST("/", controller.CreateProbe)

	}
}


func SetupCommandRoutes(router *gin.Engine, controller controller.CommandController) {
	commandGroup := router.Group("/command")
	{
		commandGroup.POST("/", controller.ExecuteCommand)
	
	}
}



