package service

import (
	"fmt"
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
	CreateProbe(probe *model.Probe) (model.Probe, error)
	ExecuteCommand(command string, id uint) (model.Probe, error)
}

type probeService struct{
	probeRepository repository.ProbeRepository
}


func (p *probeService) CreateProbe(probe *model.Probe) (model.Probe, error) {
	createdProbe, err := p.probeRepository.CreateProbe(*probe)
	if err != nil {
		return model.Probe{}, err
	}
	return createdProbe, nil
}


func (p *probeService) ExecuteCommand(command string, id uint) (model.Probe, error) {
	var probe model.Probe

	probe, err := p.probeRepository.GetProbeByID(id)
	if err != nil {
		return model.Probe{}, fmt.Errorf("probe not found: %v", err)
	}

	movementProbe(&probe, command)

	probe, err = p.probeRepository.UpdateProbe(&probe)
	if err != nil {
		return model.Probe{}, fmt.Errorf("failed to save probe state: %v", err)
	}

	return probe, nil
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