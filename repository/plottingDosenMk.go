package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PlottingDosenMkRepository interface {
	GetPlottingDosenMk() ([]model.PlottingDosenMk, error)
}

type plottingDosenMkRepository struct {
	db *gorm.DB
}

func NewPlottingDosenMkRepo(db *gorm.DB) PlottingDosenMkRepository {
	return &plottingDosenMkRepository{db}
}

func (p *plottingDosenMkRepository) GetPlottingDosenMk() ([]model.PlottingDosenMk, error) {
	var plottingDosenMk []model.PlottingDosenMk
	err := p.db.Find(&plottingDosenMk).Error
	if err != nil {
		return []model.PlottingDosenMk{}, err
	}
	return plottingDosenMk, nil
}
