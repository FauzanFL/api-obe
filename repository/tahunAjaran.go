package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type TahunAjaranRepository interface {
	GetTahunAjaran() ([]model.TahunAjaran, error)
}

type tahunAjaranRepository struct {
	db *gorm.DB
}

func NewTahunAjaranRepo(db *gorm.DB) TahunAjaranRepository {
	return &tahunAjaranRepository{db}
}

func (p *tahunAjaranRepository) GetTahunAjaran() ([]model.TahunAjaran, error) {
	var tahunAjaran []model.TahunAjaran
	err := p.db.Find(&tahunAjaran).Error
	if err != nil {
		return []model.TahunAjaran{}, err
	}
	return tahunAjaran, nil
}
