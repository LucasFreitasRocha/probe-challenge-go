package service

import (
	"regexp"

	"github.com/LucasFreitasRocha/probe-challenge-go/config/logger"
	"github.com/LucasFreitasRocha/probe-challenge-go/config/rest_err"
	"github.com/LucasFreitasRocha/probe-challenge-go/model"
	"github.com/LucasFreitasRocha/probe-challenge-go/repository"
)

var spinLeft = map[string]string{
	"N": "W",
	"E": "N",
	"S": "E",
	"W": "S",
}

var spinRight = map[string]string{
	"N": "E",
	"E": "S",
	"S": "W",
	"W": "N",
}

var directions = []string{"N", "E", "S", "W"}
var validCommands = []string{"L", "R", "M"}

func NewProbeService(
	probeRepository repository.ProbeRepository,
) ProbeService {
	logger.Info("Creating new probe service")
	return &probeService{
		probeRepository: probeRepository,
	}
}


type ProbeService interface {
	CreateProbe(probe *model.Probe) (model.Probe, *rest_err.RestErr)
	ExecuteCommand(command string, id uint) (model.Probe, *rest_err.RestErr)
}

type probeService struct{
	probeRepository repository.ProbeRepository
}


func (p *probeService) CreateProbe(probe *model.Probe) (model.Probe, *rest_err.RestErr) {
	probeResponse, err := p.probeRepository.GetProbeBy(probe.Name, probe.PositionX, probe.PositionY)
	if( err != nil) {
		rest_err.NewInternalServerError(err.Error())
	}
	err = validateCreateProbe(probe, probeResponse); if err != nil {
		logger.Error("Probe validation failed", err)
		return model.Probe{}, err
	}	
	return  p.probeRepository.CreateProbe(*probe)
}


func (p *probeService) ExecuteCommand(command string, id uint) (model.Probe, *rest_err.RestErr) {
	var probe model.Probe

	probe, err := p.probeRepository.GetProbeByID(id)
	if err != nil {
		return model.Probe{}, err
	}

	if containsInvalidChars(command) {
		logger.Error("Invalid characters in command", nil)	
		return model.Probe{}, rest_err.NewBadRequestError("command contains invalid characters, only L, R and M are allowed")
	}
	movementProbe(&probe, command)
	return p.probeRepository.UpdateProbe(&probe)
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



func containsInvalidChars(s string) bool {
	re := regexp.MustCompile(`[^LRM]`)
	return re.MatchString(s)
}

func contains(strs []string, target string) bool {
	for _, s := range strs {
			if s == target {
					return true
			}
	}
	return false
}


func validateCreateProbe(probe, probeResponse *model.Probe) *rest_err.RestErr {
 	causes := []rest_err.Causes{}

	if !contains(directions, probe.Direction) {
		causes = append(causes, rest_err.Causes{
			Field:   "direction",
			Message: "direction not valid, must be one of: N, E, S, W",
		})	
	}
	if probe.Name == ""  {
		causes = append(causes, rest_err.Causes{
			Field:   "name",
			Message: "probe name cannot be empty",
		})
		
	}

	if probeResponse.Name == probe.Name {
		causes = append(causes, rest_err.Causes{
			Field:   "name",
			Message: "probe already exists with this name",
		})	
	}

	if probeResponse.PositionX == probe.PositionX && probeResponse.PositionY == probe.PositionY {
		causes = append(causes, rest_err.Causes{
			Field:   "position",	
			Message: "probe already exists with this position",
		})
	}
	if len(causes) > 0 {
		return rest_err.NewBadRequestValidationError(
			"invalid probe data",
			causes,
		)
	}
	return nil
}