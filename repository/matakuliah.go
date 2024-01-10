package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type MataKuliahRepository interface {
	GetMataKuliah() ([]model.MataKuliah, error)
	GetMataKuliahById(id int) (model.MataKuliah, error)
	GetCLOFromMataKuliah(id int) ([]model.CLO, error)
	GetPLOFromMataKuliah(id int) ([]model.PLO, error)
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

func (m *mataKuliahRepository) GetCLOFromMataKuliah(id int) ([]model.CLO, error) {
	var clo []model.CLO
	err := m.dbKurikulum.Model(&model.CLO{}).Where("mk_id =?", id).Find(&clo).Error
	if err != nil {
		return []model.CLO{}, err
	}
	return clo, nil
}

func (m *mataKuliahRepository) GetPLOFromMataKuliah(id int) ([]model.PLO, error) {
	var plo []model.PLO
	err := m.dbKurikulum.Model(&model.PLO{}).Where("plo_id IN ?", m.dbKurikulum.Table("plo_clo").Where("clo_id IN ?", m.dbKurikulum.Table("clo").Where("mkd_id = ?", id).Select("id")).Select("plo_id")).Find(&plo).Error
	if err != nil {
		return []model.PLO{}, err
	}
	return plo, nil
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
