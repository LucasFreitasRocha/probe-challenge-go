package routes

import (
	"probe-challenge/controller"
	"github.com/gin-gonic/gin"
)

func SetupProbeRoutes(router *gin.Engine) {
	
	probeGroup := router.Group("/probes")
	{
		probeGroup.POST("/", controller.CreateProbe)

		probeGroup.POST("/:id/command", controller.ExecuteCommand)
		
	}
}