package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type KelasRepository interface {
	GetKelas() ([]model.Kelas, error)
}

type kelasRepository struct {
	db *gorm.DB
}

func NewKelasRepo(db *gorm.DB) KelasRepository {
	return &kelasRepository{db}
}

func (k *kelasRepository) GetKelas() ([]model.Kelas, error) {
	var kelas []model.Kelas
	err := k.db.Find(&kelas).Error
	if err != nil {
		return []model.Kelas{}, err
	}
	return kelas, nil
}
