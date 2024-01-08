package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type KelasRepository interface {
	GetKelas() ([]model.Kelas, error)
}

type kelasRepository struct {
	dbPenilaian *gorm.DB
}

func NewKelasRepo(dbPenilaian *gorm.DB) KelasRepository {
	return &kelasRepository{dbPenilaian}
}

func (k *kelasRepository) GetKelas() ([]model.Kelas, error) {
	var kelas []model.Kelas
	err := k.dbPenilaian.Find(&kelas).Error
	if err != nil {
		return []model.Kelas{}, err
	}
	return kelas, nil
}
