package service

import (
	"fmt"
	"probe-challenge/database"
	"probe-challenge/model"
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

type ProbeServiceInterface interface {
	CreateProbe(probe *model.Probe) (model.Probe, error)
	ExecuteCommand( command string, id uint) (model.Probe, error)
}

type ProbeService struct{}

var ProbeServiceSingleton *ProbeService; 

func (p ProbeService) CreateProbe(probe *model.Probe) (model.Probe, error) {
	if err := database.DB.Create(probe).Error; err != nil {
		return model.Probe{}, err
	}
	return *probe, nil
}

// movementProbe updates the probe's position or direction based on the command.





func (p *ProbeService) GetProbeService() *ProbeService{
	if ProbeServiceSingleton == nil {
		fmt.Println("Initializing ProbeServiceSingleton")
		ProbeServiceSingleton = &ProbeService{}
	}
	return ProbeServiceSingleton;
}


func (p *ProbeService) ExecuteCommand(command string, id uint) (model.Probe, error) {
	probe := &model.Probe{}
	if err := database.DB.First(probe, id).Error; err != nil {
		return model.Probe{}, fmt.Errorf("probe not found: %v", err)
	}

	movementProbe(probe, command)

	if err := database.DB.Save(probe).Error; err != nil {
		return model.Probe{}, fmt.Errorf("failed to save probe state: %v", err)
	}

	return *probe, nil
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