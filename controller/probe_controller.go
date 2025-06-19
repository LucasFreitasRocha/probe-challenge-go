package controller

import (
	"strconv"

	"github.com/LucasFreitasRocha/probe-challenge-go/dto"
	"github.com/LucasFreitasRocha/probe-challenge-go/model"
	"github.com/LucasFreitasRocha/probe-challenge-go/service"
	"github.com/gin-gonic/gin"
)

func NewProbeController(
	probeService service.ProbeService,
) ProbeController {
	return &probeController{
		probeService: probeService,
	}
}


type ProbeController interface {
	CreateProbe(c *gin.Context)
	ExecuteCommand(c *gin.Context)
}

type probeController struct {
	probeService service.ProbeService
}

func (pc *probeController) CreateProbe(c *gin.Context) {
	var probe model.Probe
	if err := c.ShouldBindJSON(&probe); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	probe , err := pc.probeService.CreateProbe(&probe)
	

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	var probeDto = *dto.FromModel(&probe)
	c.JSON(200, probeDto)
}

func (pc *probeController) ExecuteCommand(c *gin.Context) {
	id := c.Param("id")
	var command dto.CommandDTO
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid command data",
		})
		return
	}
	probeID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid probe ID",
		})
		return
	}
	result, err := pc.probeService.ExecuteCommand(command.Command, uint(probeID))

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	var probeDto = *dto.FromModel(&result)
	c.JSON(200, probeDto)
}


