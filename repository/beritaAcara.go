package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type BeritaAcaraRepository interface {
	GetBeritaAcara() ([]model.BeritaAcara, error)
}

type beritaAcaraRepository struct {
	db *gorm.DB
}

func NewBeritaAcaraRepo(db *gorm.DB) BeritaAcaraRepository {
	return &beritaAcaraRepository{db}
}

func (b *beritaAcaraRepository) GetBeritaAcara() ([]model.BeritaAcara, error) {
	var beritaAcara []model.BeritaAcara
	err := b.db.Find(&beritaAcara).Error
	if err != nil {
		return []model.BeritaAcara{}, err
	}
	return beritaAcara, nil
}
