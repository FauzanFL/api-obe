package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type CloRepository interface {
	GetClo() ([]model.CLO, error)
}

type cloRepository struct {
	db *gorm.DB
}

func NewCloRepo(db *gorm.DB) CloRepository {
	return &cloRepository{db}
}

func (c *cloRepository) GetClo() ([]model.CLO, error) {
	var clo []model.CLO
	err := c.db.Find(&clo).Error
	if err != nil {
		return []model.CLO{}, err
	}
	return clo, nil
}
