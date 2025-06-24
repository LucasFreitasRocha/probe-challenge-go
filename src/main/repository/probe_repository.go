package repository

import (
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config/logger"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/config/rest_err"
	"github.com/LucasFreitasRocha/probe-challenge-go/src/main/model"
	"gorm.io/gorm"
)

func NewProbeRepository(database *gorm.DB) ProbeRepository {
	logger.Info("Creating new probe repository")
	return &probeRepository{
		database: database,
	}
}

type probeRepository struct {
	database *gorm.DB
}

// GetProbeBy implements ProbeRepository.


// GetProbeByName implements ProbeRepository.

type ProbeRepository interface {
	CreateProbe(probe model.Probe) (model.Probe, *rest_err.RestErr)
	UpdateProbe(probe *model.Probe) (model.Probe, *rest_err.RestErr)
	GetProbeByID(id uint) (model.Probe, *rest_err.RestErr)
	GetProbeBy(name string, positionX, positionY int) (*model.Probe, *rest_err.RestErr)
}

func (r *probeRepository) CreateProbe(probe model.Probe) (model.Probe, *rest_err.RestErr) {
	if err := r.database.Create(&probe).Error; err != nil {
		return model.Probe{}, rest_err.NewInternalServerError(err.Error())
	}
	return probe, nil
}

func (r *probeRepository) UpdateProbe(probe *model.Probe) (model.Probe, *rest_err.RestErr) {
	if err := r.database.Save(probe).Error; err != nil {
		return *probe, rest_err.NewInternalServerError(err.Error())
	}
	return *probe, nil
}

func (r *probeRepository) GetProbeByID(id uint) (model.Probe, *rest_err.RestErr) {
	var probe model.Probe
	if err := r.database.First(&probe, id).Error; err != nil {
		return model.Probe{}, rest_err.NewNotFoundError(err.Error())
	}
	return probe, nil
}

func (r *probeRepository) GetProbeBy(name string, positionX int, positionY int) (*model.Probe, *rest_err.RestErr) {
	var probe = &model.Probe{}
	if err := r.database.Where("name = ? OR (position_x = ? AND position_y = ?)", name, positionX, positionY).First(&probe).Error; err != nil {
		return probe, rest_err.NewNotFoundError(err.Error())
	}
	return probe, nil
}