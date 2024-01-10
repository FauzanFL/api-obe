package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PenilaianRepository interface {
	GetPenilaian() ([]model.Penilaian, error)
	GetPenilaianById(id int) (model.Penilaian, error)
	GetPenilaianByKelas(kelasId int) ([]model.Penilaian, error)
	CreatePenilaian(penilaian model.Penilaian) error
	UpdatePenilaian(penilaian model.Penilaian) error
	DeletePenilaian(id int) error
}

type penilaianRepository struct {
	dbPenilaian *gorm.DB
}

func NewPenilaianRepo(dbPenilaian *gorm.DB) PenilaianRepository {
	return &penilaianRepository{dbPenilaian}
}

func (p *penilaianRepository) GetPenilaian() ([]model.Penilaian, error) {
	var penilaian []model.Penilaian
	err := p.dbPenilaian.Find(&penilaian).Error
	if err != nil {
		return []model.Penilaian{}, err
	}
	return penilaian, nil
}

func (p *penilaianRepository) GetPenilaianById(id int) (model.Penilaian, error) {
	var penilaian model.Penilaian
	err := p.dbPenilaian.First(&penilaian, id).Error
	if err != nil {
		return model.Penilaian{}, err
	}
	return penilaian, nil
}

func (p *penilaianRepository) GetPenilaianByKelas(kelasId int) ([]model.Penilaian, error) {
	var penilaian []model.Penilaian
	err := p.dbPenilaian.Where("mhs_id IN ?", p.dbPenilaian.Table("mahasiswa").Where("kelas_id = ?", kelasId).Select("id")).Find(&penilaian).Error
	if err != nil {
		return []model.Penilaian{}, err
	}
	return penilaian, nil
}

func (p *penilaianRepository) CreatePenilaian(penilaian model.Penilaian) error {
	err := p.dbPenilaian.Create(&penilaian).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *penilaianRepository) UpdatePenilaian(penilaian model.Penilaian) error {
	err := p.dbPenilaian.Save(&penilaian).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *penilaianRepository) DeletePenilaian(id int) error {
	err := p.dbPenilaian.Delete(&model.Penilaian{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
