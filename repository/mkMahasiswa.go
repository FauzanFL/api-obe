package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type MkMahasiswaRepository interface {
	GetMkMahasiswa() ([]model.MkMahasiswa, error)
}

type mkMahasiswaRepository struct {
	db *gorm.DB
}

func NewMkMahasiswaRepo(db *gorm.DB) MkMahasiswaRepository {
	return &mkMahasiswaRepository{db}
}

func (m *mkMahasiswaRepository) GetMkMahasiswa() ([]model.MkMahasiswa, error) {
	var mkMahasiswa []model.MkMahasiswa
	err := m.db.Find(&mkMahasiswa).Error
	if err != nil {
		return []model.MkMahasiswa{}, err
	}
	return mkMahasiswa, nil
}
