package controller

import (
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config/logger"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/dto"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/service"
	"github.com/gin-gonic/gin"
)

func NewCommandController(
	commandService service.CommandService,
) CommandController {
	logger.Info("Creating new command controller")
	return &commandController{
		commandService: commandService,
	}
}

type CommandController interface {
	ExecuteCommand(co *gin.Context)
}
type commandController struct {
	commandService service.CommandService
}


func (c *commandController) ExecuteCommand(co *gin.Context) {
	var command dto.CommandDTO
	if err := co.ShouldBindJSON(&command); err != nil {
		notValidPayload(co)
		return
	}
	probe, err := c.commandService.ExecuteCommand(command.Command, command.IdProbe)
	if err != nil {
		logger.Error("Failed to execute command", err)
		 co.JSON(err.Code, err)
		 return
	}
	var probeDto = *dto.FromModel(&probe)
	co.JSON(200, probeDto)
}
