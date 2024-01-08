package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PerancanganObeRepository interface {
	GetPerancanganObe() ([]model.PerancanganObe, error)
	GetPerancanganObeById(id int) (model.PerancanganObe, error)
	CreatePerancanganObe(perancanganObe model.PerancanganObe) error
	UpdatePerancanganObe(perancanganObe model.PerancanganObe) error
	DeletePerancanganObe(id int) error
}

type perancanganObeRepository struct {
	dbKurikulum *gorm.DB
}

func NewPerancanganObeRepo(dbKurikulum *gorm.DB) PerancanganObeRepository {
	return &perancanganObeRepository{dbKurikulum}
}

func (p *perancanganObeRepository) GetPerancanganObe() ([]model.PerancanganObe, error) {
	var perancanganObe []model.PerancanganObe
	err := p.dbKurikulum.Find(&perancanganObe).Error
	if err != nil {
		return []model.PerancanganObe{}, err
	}
	return perancanganObe, nil
}

func (p *perancanganObeRepository) GetPerancanganObeById(id int) (model.PerancanganObe, error) {
	var perancanganObe model.PerancanganObe
	err := p.dbKurikulum.First(&perancanganObe, id).Error
	if err != nil {
		return model.PerancanganObe{}, err
	}
	return perancanganObe, nil
}

func (p *perancanganObeRepository) CreatePerancanganObe(perancanganObe model.PerancanganObe) error {
	err := p.dbKurikulum.Create(&perancanganObe).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *perancanganObeRepository) UpdatePerancanganObe(perancanganObe model.PerancanganObe) error {
	err := p.dbKurikulum.Save(&perancanganObe).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *perancanganObeRepository) DeletePerancanganObe(id int) error {
	err := p.dbKurikulum.Delete(&model.PerancanganObe{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
