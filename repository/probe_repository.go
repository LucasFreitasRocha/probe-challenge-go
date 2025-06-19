package repository

import (
	"github.com/LucasFreitasRocha/probe-challenge-go/model"
	"github.com/LucasFreitasRocha/probe-challenge-go/config/rest_err"
	"gorm.io/gorm"
)


func NewProbeRepository(database *gorm.DB) ProbeRepository {
	return &probeRepository{
		database: database,
	}
}


type probeRepository struct {
	database *gorm.DB
}

type ProbeRepository interface {
	CreateProbe(probe model.Probe) (model.Probe, *rest_err.RestErr)
	UpdateProbe(probe *model.Probe) (model.Probe, *rest_err.RestErr)
	GetProbeByID(id uint) (model.Probe, *rest_err.RestErr)
}


func (r *probeRepository) CreateProbe(probe model.Probe) (model.Probe, *rest_err.RestErr) {
	if err := r.database.Create(&probe).Error; err != nil {
		return probe, rest_err.NewInternalServerError("error creating probe")	
	}
	return probe, nil
}


func (r *probeRepository) UpdateProbe(probe *model.Probe) (model.Probe, *rest_err.RestErr) {
	if err := r.database.Save(probe).Error; err != nil {
		return *probe, rest_err.NewInternalServerError("error updating probe")
	}
	return *probe, nil
}


func (r *probeRepository) GetProbeByID(id uint) (model.Probe, *rest_err.RestErr) {
	var probe model.Probe
	if err := r.database.First(&probe, id).Error; err != nil {
		return model.Probe{}, rest_err.NewBadRequestError("probe not found: " + err.Error()) 
	}
	return probe, nil
}