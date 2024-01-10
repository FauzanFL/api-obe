package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PloCloRepository interface {
	Get() ([]model.PLO_CLO, error)
	Create(ploClo model.PLO_CLO) error
	Delete(id int) error
}

type ploCloRepository struct {
	dbKurikulum *gorm.DB
}

func NewPloCloRepo(dbKurikulum *gorm.DB) PloCloRepository {
	return &ploCloRepository{dbKurikulum}
}

func (p *ploCloRepository) Get() ([]model.PLO_CLO, error) {
	var result []model.PLO_CLO
	err := p.dbKurikulum.Find(&result).Error
	if err != nil {
		return []model.PLO_CLO{}, err
	}
	return result, nil
}

func (p *ploCloRepository) Create(ploClo model.PLO_CLO) error {
	err := p.dbKurikulum.Create(&ploClo).Error
	return err
}

func (p *ploCloRepository) Delete(id int) error {
	err := p.dbKurikulum.Delete(&model.PLO_CLO{}, id).Error
	return err
}
