package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type DosenRepository interface {
	GetDosen() ([]model.Dosen, error)
}

type dosenRepository struct {
	dbUser *gorm.DB
}

func NewDosenRepo(dbUser *gorm.DB) DosenRepository {
	return &dosenRepository{dbUser}
}

func (d *dosenRepository) GetDosen() ([]model.Dosen, error) {
	var dosen []model.Dosen
	err := d.dbUser.Find(&dosen).Error
	if err != nil {
		return []model.Dosen{}, err
	}
	return dosen, nil
}
