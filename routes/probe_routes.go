package routes

import (
	"github.com/LucasFreitasRocha/probe-challenge-go/controller"
	"github.com/gin-gonic/gin"
)

func SetupProbeRoutes(router *gin.Engine, controller controller.ProbeController) {

	probeGroup := router.Group("/probes")
	{
		probeGroup.POST("/", controller.CreateProbe)

		probeGroup.POST("/:id/command", controller.ExecuteCommand)

	}
}



