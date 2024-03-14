package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type KelasRepository interface {
	GetKelas() ([]model.Kelas, error)
	GetKelasByKeyword(keyword string) ([]model.Kelas, error)
	GetKelasById(id int) (model.Kelas, error)
	CreateKelas(kelas model.Kelas) error
	UpdateKelas(kelas model.Kelas) error
	DeleteKelas(id int) error
}

type kelasRepository struct {
	dbPenilaian *gorm.DB
}

func NewKelasRepo(dbPenilaian *gorm.DB) KelasRepository {
	return &kelasRepository{dbPenilaian}
}

func (k *kelasRepository) GetKelas() ([]model.Kelas, error) {
	var kelas []model.Kelas
	err := k.dbPenilaian.Find(&kelas).Error
	if err != nil {
		return []model.Kelas{}, err
	}
	return kelas, nil
}

func (k *kelasRepository) GetKelasByKeyword(keyword string) ([]model.Kelas, error) {
	var kelas []model.Kelas
	err := k.dbPenilaian.Where("kode_kelas LIKE ?", "%"+keyword+"%").Find(&kelas).Error
	if err != nil {
		return []model.Kelas{}, err
	}
	return kelas, nil
}

func (k *kelasRepository) GetKelasById(id int) (model.Kelas, error) {
	var kelas model.Kelas
	err := k.dbPenilaian.Where("id = ?", id).First(&kelas).Error
	if err != nil {
		return model.Kelas{}, err
	}
	return kelas, nil
}

func (k *kelasRepository) CreateKelas(kelas model.Kelas) error {
	err := k.dbPenilaian.Create(&kelas).Error
	return err
}

func (k *kelasRepository) UpdateKelas(kelas model.Kelas) error {
	err := k.dbPenilaian.Save(&kelas).Error
	return err
}

func (k *kelasRepository) DeleteKelas(id int) error {
	err := k.dbPenilaian.Delete(&model.Kelas{}, id).Error
	return err
}
