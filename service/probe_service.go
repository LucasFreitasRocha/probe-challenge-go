package service

import (

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
	GetProbeByID(id uint) (model.Probe, *rest_err.RestErr)
	UpdateProbe(probe *model.Probe) (model.Probe, *rest_err.RestErr)
}

type probeService struct {
	probeRepository repository.ProbeRepository
}



func (p *probeService) CreateProbe(probe *model.Probe) (model.Probe, *rest_err.RestErr) {
	probeResponse, err := p.probeRepository.GetProbeBy(probe.Name, probe.PositionX, probe.PositionY)
	if err != nil {
		rest_err.NewInternalServerError(err.Error())
	}
	err = validateCreateProbe(probe, probeResponse)
	if err != nil {
		logger.Error("Probe validation failed", err)
		return model.Probe{}, err
	}
	return p.probeRepository.CreateProbe(*probe)
}


func (p *probeService) GetProbeByID(id uint) (model.Probe, *rest_err.RestErr) {
	return p.probeRepository.GetProbeByID(id)
}


func (p *probeService) UpdateProbe(probe *model.Probe) (model.Probe, *rest_err.RestErr) {
	return p.probeRepository.UpdateProbe(probe)
}




func validateCreateProbe(probe, probeResponse *model.Probe) *rest_err.RestErr {
	causes := []rest_err.Causes{}

	if validateString(probe.Direction, `[^NESW]$`) {
		causes = append(causes, rest_err.Causes{
			Field:   "direction",
			Message: "direction not valid, must be one of: N, E, S, W",
		})
	}
	if probe.Name == "" {
		causes = append(causes, rest_err.Causes{
			Field:   "name",
			Message: "probe name cannot be empty",
		})

	}

	if probeResponse.Name == probe.Name {
		causes = append(causes, rest_err.Causes{
			Field:   "name",
			Message: "A probe already exists with this name",
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
