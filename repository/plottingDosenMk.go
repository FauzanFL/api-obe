package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PlottingDosenMkRepository interface {
	GetPlottingDosenMk() ([]model.PlottingDosenMk, error)
	GetPlottingDosenByMkId(mkId int) ([]model.PlottingDosenMk, error)
	CreatePlottingDosenMk(plottingDosenMk model.PlottingDosenMk) error
	DeletePlottingDosenMk(id int) error
}

type plottingDosenMkRepository struct {
	dbKurikulum *gorm.DB
}

func NewPlottingDosenMkRepo(dbKurikulum *gorm.DB) PlottingDosenMkRepository {
	return &plottingDosenMkRepository{dbKurikulum}
}

func (p *plottingDosenMkRepository) GetPlottingDosenMk() ([]model.PlottingDosenMk, error) {
	var plottingDosenMk []model.PlottingDosenMk
	err := p.dbKurikulum.Find(&plottingDosenMk).Error
	if err != nil {
		return []model.PlottingDosenMk{}, err
	}
	return plottingDosenMk, nil
}

func (p *plottingDosenMkRepository) GetPlottingDosenByMkId(mkId int) ([]model.PlottingDosenMk, error) {
	var plottingDosenMk []model.PlottingDosenMk
	err := p.dbKurikulum.Where("mkId =?", mkId).Find(&plottingDosenMk).Error
	if err != nil {
		return []model.PlottingDosenMk{}, err
	}
	return plottingDosenMk, nil
}

func (p *plottingDosenMkRepository) CreatePlottingDosenMk(plottingDosenMk model.PlottingDosenMk) error {
	err := p.dbKurikulum.Create(&plottingDosenMk).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *plottingDosenMkRepository) DeletePlottingDosenMk(id int) error {
	err := p.dbKurikulum.Delete(&model.PlottingDosenMk{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
