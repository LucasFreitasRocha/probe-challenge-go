package service

import (
	"regexp"

	"github.com/LucasFreitasRocha/probe-challenge-go/config/logger"
	"github.com/LucasFreitasRocha/probe-challenge-go/config/rest_err"
	"github.com/LucasFreitasRocha/probe-challenge-go/model"
)

type CommandService interface {
	ExecuteCommand(command string, id uint) (model.Probe, *rest_err.RestErr)
}

type commandService struct {
	probeService ProbeService
}

// ExecuteCommand implements CommandService.
func (c *commandService) ExecuteCommand(command string, id uint) (model.Probe, *rest_err.RestErr) {
	var probe model.Probe
	probe, err := c.probeService.GetProbeByID(id)
	if err != nil {
		return model.Probe{}, err
	}

	if validateString(command, `[^LRM]`) {
		logger.Error("Invalid characters in command", nil)
		return model.Probe{}, rest_err.NewBadRequestError("command contains invalid characters, only L, R and M are allowed")
	}
	movementProbe(&probe, command)
	return c.probeService.UpdateProbe(&probe)
}

func NewCommandService(probeService ProbeService) CommandService {
	return &commandService{ probeService: probeService}
}




func movementProbe(probe *model.Probe, commands string) {
	for _, command := range commands {
		switch string(command) {
		case "L":
			probe.Direction = spinLeft[probe.Direction]
		case "R":
			probe.Direction = spinRight[probe.Direction]
		case "M":
			switch probe.Direction {
			case "N":
				probe.PositionY++
			case "S":
				probe.PositionY--
			case "E":
				probe.PositionX++
			case "W":
				probe.PositionX--
			}
		}
	}
}


func validateString(s string, r string) bool {
	re := regexp.MustCompile(r)
	return re.MatchString(s)
}