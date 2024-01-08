package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type DosenRepository interface {
	GetDosen() ([]model.Dosen, error)
}

type dosenRepository struct {
	db *gorm.DB
}

func NewDosenRepo(db *gorm.DB) DosenRepository {
	return &dosenRepository{db}
}

func (d *dosenRepository) GetDosen() ([]model.Dosen, error) {
	var dosen []model.Dosen
	err := d.db.Find(&dosen).Error
	if err != nil {
		return []model.Dosen{}, err
	}
	return dosen, nil
}
