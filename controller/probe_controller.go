package controller

import (
	"probe-challenge/dto"
	"probe-challenge/model"
	"probe-challenge/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateProbe(c *gin.Context) {
	var probe model.Probe
	if err := c.ShouldBindJSON(&probe); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	probeService := service.ProbeServiceSingleton.GetProbeService()
	probe, err := probeService.CreateProbe(&probe)
	

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var probeDto = *dto.FromModel(&probe)
	c.JSON(200, probeDto)
}

func ExecuteCommand(c *gin.Context) {
	id := c.Param("id")
	var command dto.CommandDTO
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid command data",
		})
		return
	}

	probeService := service.ProbeServiceSingleton.GetProbeService()
	probeID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid probe ID",
		})
		return
	}
	result, err := probeService.ExecuteCommand(command.Command, uint(probeID))

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	var probeDto = *dto.FromModel(&result)
	c.JSON(200, probeDto)
}


