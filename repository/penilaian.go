package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PenilaianRepository interface {
	GetPenilaian() ([]model.Penilaian, error)
	GetPenilaianById(id int) (model.Penilaian, error)
	GetPenilaianByKelasIdAndMkId(kelasId int, mkId int) (model.Penilaian, error)
	CreatePenilaian(penilaian model.Penilaian) error
	UpdatePenilaian(penilaian model.Penilaian) error
	UpdateStatusToFinal(id int) error
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

func (p *penilaianRepository) GetPenilaianByKelasIdAndMkId(kelasId int, mkId int) (model.Penilaian, error) {
	var penilaian model.Penilaian
	err := p.dbPenilaian.Where("kelas_id = ? AND mk_id = ?", kelasId, mkId).First(&penilaian).Error
	if err != nil {
		return model.Penilaian{}, err
	}
	return penilaian, nil
}

func (p *penilaianRepository) UpdateStatusToFinal(id int) error {
	err := p.dbPenilaian.Model(&model.Penilaian{}).Where("id = ?", id).Update("status", "final").Error
	if err != nil {
		return err
	}
	return nil
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
