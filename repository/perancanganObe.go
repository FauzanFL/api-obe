package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PerancanganObeRepository interface {
	GetPerancanganObe() ([]model.PerancanganObe, error)
}

type perancanganObeRepository struct {
	db *gorm.DB
}

func NewPerancanganObeRepo(db *gorm.DB) PerancanganObeRepository {
	return &perancanganObeRepository{db}
}

func (m *perancanganObeRepository) GetPerancanganObe() ([]model.PerancanganObe, error) {
	var perancanganObe []model.PerancanganObe
	err := m.db.Find(&perancanganObe).Error
	if err != nil {
		return []model.PerancanganObe{}, err
	}
	return perancanganObe, nil
}
