package controller

import (
	"probe-challenge/dto"
	"probe-challenge/model"
	"probe-challenge/service"

	"github.com/gin-gonic/gin"
)

func CreateProbe(c *gin.Context) {
	var probe model.Probe
	if err := c.ShouldBindJSON(&probe); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	probe, err := service.CreateProbe(&probe)
	

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var probeDto = *dto.FromModel(&probe)
	c.JSON(200, probeDto)
}


