package service
import (
	"probe-challenge/model"
	"probe-challenge/database"
)

func CreateProbe(probe *model.Probe) (model.Probe, error) {
	if err := database.DB.Create(probe).Error; err != nil {
		return model.Probe{}, err
	}
	return *probe, nil
}