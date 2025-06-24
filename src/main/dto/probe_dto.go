package dto

import "github.com/LucasFreitasRocha/probe-challenge-go/src/main/model"

type ProbeDTO struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	PositionX int    `json:"position_x"`
	PositionY int    `json:"position_y"`
	Direction string `json:"direction"`
}

func (p *ProbeDTO) ToModel() *model.Probe {
	return &model.Probe{
		Id:        p.Id,
		Name:      p.Name,
		PositionX: p.PositionX,
		PositionY: p.PositionY,
		Direction: p.Direction,
	}
}

func FromModel(model *model.Probe) *ProbeDTO {
	return &ProbeDTO{
		Id:        model.Id,
		Name:      model.Name,
		PositionX: model.PositionX,
		PositionY: model.PositionY,
		Direction: model.Direction,
	}
}
