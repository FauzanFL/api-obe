package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	GetMahasiswa() ([]model.Mahasiswa, error)
}

type mahasiswaRepository struct {
	db *gorm.DB
}

func NewMahasiswaRepo(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepository{db}
}

func (m *mahasiswaRepository) GetMahasiswa() ([]model.Mahasiswa, error) {
	var mahasiswa []model.Mahasiswa
	err := m.db.Find(&mahasiswa).Error
	if err != nil {
		return []model.Mahasiswa{}, err
	}
	return mahasiswa, nil
}
