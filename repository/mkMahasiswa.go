package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type MkMahasiswaRepository interface {
	GetMkMahasiswa() ([]model.MkMahasiswa, error)
	CreateMkMahasiswa(mkMahasiswa model.MkMahasiswa) error
	DeleteMkMahasiswa(id int) error
}

type mkMahasiswaRepository struct {
	dbPenilaian *gorm.DB
}

func NewMkMahasiswaRepo(dbPenilaian *gorm.DB) MkMahasiswaRepository {
	return &mkMahasiswaRepository{dbPenilaian}
}

func (m *mkMahasiswaRepository) GetMkMahasiswa() ([]model.MkMahasiswa, error) {
	var mkMahasiswa []model.MkMahasiswa
	err := m.dbPenilaian.Find(&mkMahasiswa).Error
	if err != nil {
		return []model.MkMahasiswa{}, err
	}
	return mkMahasiswa, nil
}

func (m *mkMahasiswaRepository) CreateMkMahasiswa(mkMahasiswa model.MkMahasiswa) error {
	err := m.dbPenilaian.Create(&mkMahasiswa).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *mkMahasiswaRepository) DeleteMkMahasiswa(id int) error {
	err := m.dbPenilaian.Delete(&model.MkMahasiswa{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
