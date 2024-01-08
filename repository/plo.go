package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PloRepository interface {
	GetPlo() ([]model.PLO, error)
	GetPloById(id int) (model.PLO, error)
	CreatePlo(plo model.PLO) error
	UpdatePlo(plo model.PLO) error
	DeletePlo(id int) error
}

type ploRepository struct {
	dbKurikulum *gorm.DB
}

func NewPloRepo(dbKurikulum *gorm.DB) PloRepository {
	return &ploRepository{dbKurikulum}
}

func (p *ploRepository) GetPlo() ([]model.PLO, error) {
	var plo []model.PLO
	err := p.dbKurikulum.Find(&plo).Error
	if err != nil {
		return []model.PLO{}, err
	}
	return plo, nil
}

func (p *ploRepository) GetPloById(id int) (model.PLO, error) {
	var plo model.PLO
	err := p.dbKurikulum.First(&plo, id).Error
	if err != nil {
		return model.PLO{}, err
	}
	return plo, nil
}

func (p *ploRepository) CreatePlo(plo model.PLO) error {
	err := p.dbKurikulum.Create(&plo).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ploRepository) UpdatePlo(plo model.PLO) error {
	err := p.dbKurikulum.Save(&plo).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *ploRepository) DeletePlo(id int) error {
	err := p.dbKurikulum.Delete(&model.PLO{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
