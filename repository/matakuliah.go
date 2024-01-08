package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type MataKuliahRepository interface {
	GetMataKuliah() ([]model.MataKuliah, error)
}

type mataKuliahRepository struct {
	dbKurikulum *gorm.DB
}

func NewMataKuliahRepo(dbKurikulum *gorm.DB) MataKuliahRepository {
	return &mataKuliahRepository{dbKurikulum}
}

func (m *mataKuliahRepository) GetMataKuliah() ([]model.MataKuliah, error) {
	var mataKuliah []model.MataKuliah
	err := m.dbKurikulum.Find(&mataKuliah).Error
	if err != nil {
		return []model.MataKuliah{}, err
	}
	return mataKuliah, nil
}
