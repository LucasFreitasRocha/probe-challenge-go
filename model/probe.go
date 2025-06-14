package model

import "gorm.io/gorm"

type Probe struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"not null;unique" json:"name"`
	PositionX int    `gorm:"not null" json:"position_x"`
	PositionY int    `gorm:"not null" json:"position_y"`
	Direction string `gorm:"not null" json:"direction"`
}