package controller

import (
	"strconv"

	"github.com/LucasFreitasRocha/probe-challenge-go/config/logger"
	"github.com/LucasFreitasRocha/probe-challenge-go/config/rest_err"
	"github.com/LucasFreitasRocha/probe-challenge-go/dto"
	"github.com/LucasFreitasRocha/probe-challenge-go/model"
	"github.com/LucasFreitasRocha/probe-challenge-go/service"
	"github.com/gin-gonic/gin"
)

func NewProbeController(
	probeService service.ProbeService,
) ProbeController {
	logger.Info("Creating new probe controller")
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
		notValidPayload(c)
		return
	}
	probe , err := pc.probeService.CreateProbe(&probe)
	
	if err != nil {
		c.JSON(err.Code, gin.H{
			"error": err.Message,
			"cause": err.Causes,
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
		notValidPayload(c)
	}
	probeId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		errorMessage := rest_err.NewBadRequestError(
			"id not valid",
		)
		c.JSON(errorMessage.Code, errorMessage)
		return
	}
	result, e := pc.probeService.ExecuteCommand(command.Command, uint(probeId))
	
	if e != nil {
		c.JSON(e.Code, gin.H{
			"error": e.Message,
		})
		return
	}
	var probeDto = *dto.FromModel(&result)
	c.JSON(200, probeDto)
}


func notValidPayload(c *gin.Context) {
	errorMessage := rest_err.NewBadRequestError(
		"payload is not valid ",
	)
	c.JSON(errorMessage.Code, errorMessage)
}