package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type BeritaAcaraRepository interface {
	GetBeritaAcara() ([]model.BeritaAcara, error)
	GetBeritaAcaraById(id int) (model.BeritaAcara, error)
	GetBeritaAcaraByPenilaianId(penilaianId int) (model.BeritaAcara, error)
	CreateBeritaAcara(beritaAcara model.BeritaAcara) error
	DeleteBeritaAcara(id int) error
}

type beritaAcaraRepository struct {
	dbPenilaian *gorm.DB
}

func NewBeritaAcaraRepo(dbPenilaian *gorm.DB) BeritaAcaraRepository {
	return &beritaAcaraRepository{dbPenilaian}
}

func (b *beritaAcaraRepository) GetBeritaAcara() ([]model.BeritaAcara, error) {
	var beritaAcara []model.BeritaAcara
	err := b.dbPenilaian.Find(&beritaAcara).Error
	if err != nil {
		return []model.BeritaAcara{}, err
	}
	return beritaAcara, nil
}

func (b *beritaAcaraRepository) GetBeritaAcaraById(id int) (model.BeritaAcara, error) {
	var beritaAcara model.BeritaAcara
	err := b.dbPenilaian.First(&beritaAcara, id).Error
	if err != nil {
		return model.BeritaAcara{}, err
	}
	return beritaAcara, nil
}

func (b *beritaAcaraRepository) GetBeritaAcaraByPenilaianId(penilaianId int) (model.BeritaAcara, error) {
	var beritaAcara model.BeritaAcara
	err := b.dbPenilaian.Where("penilaian_id = ?", penilaianId).First(&beritaAcara).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.BeritaAcara{}, nil
		}
		return model.BeritaAcara{}, err
	}
	return beritaAcara, nil
}

func (b *beritaAcaraRepository) CreateBeritaAcara(beritaAcara model.BeritaAcara) error {
	err := b.dbPenilaian.Create(&beritaAcara).Error
	if err != nil {
		return err
	}
	return nil
}

func (b *beritaAcaraRepository) DeleteBeritaAcara(id int) error {
	err := b.dbPenilaian.Delete(&model.BeritaAcara{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
