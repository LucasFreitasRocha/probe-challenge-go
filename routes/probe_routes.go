package routes

import (
	"probe-challenge/controller"
	"github.com/gin-gonic/gin"
)

func SetupProbeRoutes(router *gin.Engine) {
	// Create a new group for probe routes
	probeGroup := router.Group("/probes")
	{
		// Define the route for creating a new probe
		probeGroup.POST("/", controller.CreateProbe)
		
		// Define the route for getting all probes
		// probeGroup.GET("/", controller.GetAllProbes)
		
		// Define the route for getting a specific probe by ID
		// probeGroup.GET("/:id", controller.GetProbeByID)
		
		// Define the route for updating a specific probe by ID
		// probeGroup.PUT("/:id", controller.UpdateProbeByID)
		
		// Define the route for deleting a specific probe by ID
		// probeGroup.DELETE("/:id", controller.DeleteProbeByID)
	}
}