package controller

import (

	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config/logger"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/dto"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/model"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/service"
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
		c.JSON(err.Code, err)
		return
	}
	var probeDto = *dto.FromModel(&probe)
	c.JSON(201, probeDto)
}


