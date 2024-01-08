package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PloRepository interface {
	GetPlo() ([]model.PLO, error)
}

type ploRepository struct {
	db *gorm.DB
}

func NewPloRepo(db *gorm.DB) PloRepository {
	return &ploRepository{db}
}

func (p *ploRepository) GetPlo() ([]model.PLO, error) {
	var plo []model.PLO
	err := p.db.Find(&plo).Error
	if err != nil {
		return []model.PLO{}, err
	}
	return plo, nil
}
