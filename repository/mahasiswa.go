package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	GetMahasiswa() ([]model.Mahasiswa, error)
}

type mahasiswaRepository struct {
	dbPenilaian *gorm.DB
}

func NewMahasiswaRepo(dbPenilaian *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepository{dbPenilaian}
}

func (m *mahasiswaRepository) GetMahasiswa() ([]model.Mahasiswa, error) {
	var mahasiswa []model.Mahasiswa
	err := m.dbPenilaian.Find(&mahasiswa).Error
	if err != nil {
		return []model.Mahasiswa{}, err
	}
	return mahasiswa, nil
}
