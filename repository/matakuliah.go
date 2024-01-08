package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type MataKuliahRepository interface {
	GetMataKuliah() ([]model.MataKuliah, error)
}

type mataKuliahRepository struct {
	db *gorm.DB
}

func NewMataKuliahRepo(db *gorm.DB) MataKuliahRepository {
	return &mataKuliahRepository{db}
}

func (m *mataKuliahRepository) GetMataKuliah() ([]model.MataKuliah, error) {
	var mataKuliah []model.MataKuliah
	err := m.db.Find(&mataKuliah).Error
	if err != nil {
		return []model.MataKuliah{}, err
	}
	return mataKuliah, nil
}
