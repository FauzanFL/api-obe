package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PenilaianRepository interface {
	GetPenilaian() ([]model.Penilaian, error)
}

type penilaianRepository struct {
	db *gorm.DB
}

func NewPenilaianRepo(db *gorm.DB) PenilaianRepository {
	return &penilaianRepository{db}
}

func (m *penilaianRepository) GetPenilaian() ([]model.Penilaian, error) {
	var penilaian []model.Penilaian
	err := m.db.Find(&penilaian).Error
	if err != nil {
		return []model.Penilaian{}, err
	}
	return penilaian, nil
}
