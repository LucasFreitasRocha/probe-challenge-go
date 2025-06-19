package service

import (
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
	createdProbe, err := p.probeRepository.CreateProbe(*probe)
	if err != nil {
		return model.Probe{}, err
	}
	return createdProbe, nil
}


func (p *probeService) ExecuteCommand(command string, id uint) (model.Probe, *rest_err.RestErr) {
	var probe model.Probe

	probe, err := p.probeRepository.GetProbeByID(id)
	if err != nil {
		return model.Probe{}, err
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