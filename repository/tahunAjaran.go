package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type TahunAjaranRepository interface {
	GetTahunAjaran() ([]model.TahunAjaran, error)
}

type tahunAjaranRepository struct {
	dbPenilaian *gorm.DB
}

func NewTahunAjaranRepo(dbPenilaian *gorm.DB) TahunAjaranRepository {
	return &tahunAjaranRepository{dbPenilaian}
}

func (p *tahunAjaranRepository) GetTahunAjaran() ([]model.TahunAjaran, error) {
	var tahunAjaran []model.TahunAjaran
	err := p.dbPenilaian.Find(&tahunAjaran).Error
	if err != nil {
		return []model.TahunAjaran{}, err
	}
	return tahunAjaran, nil
}
