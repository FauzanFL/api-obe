package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type MataKuliahRepository interface {
	GetMataKuliah() ([]model.MataKuliah, error)
	GetMataKuliahById(id int) (model.MataKuliah, error)
	CreateMataKuliah(mataKuliah model.MataKuliah) error
	UpdateMataKuliah(mataKuliah model.MataKuliah) error
	DeleteMataKuliah(id int) error
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

func (m *mataKuliahRepository) GetMataKuliahById(id int) (model.MataKuliah, error) {
	var mataKuliah model.MataKuliah
	err := m.dbKurikulum.First(&mataKuliah, id).Error
	if err != nil {
		return model.MataKuliah{}, err
	}
	return mataKuliah, nil
}

func (m *mataKuliahRepository) CreateMataKuliah(mataKuliah model.MataKuliah) error {
	err := m.dbKurikulum.Create(&mataKuliah).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *mataKuliahRepository) UpdateMataKuliah(mataKuliah model.MataKuliah) error {
	err := m.dbKurikulum.Save(&mataKuliah).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *mataKuliahRepository) DeleteMataKuliah(id int) error {
	err := m.dbKurikulum.Delete(&model.MataKuliah{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
