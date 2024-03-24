package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type IndexPenilaianRepository interface {
	GetIndexPenilaian() ([]model.IndexPenilaian, error)
	GetIndexPenilaianByNilai(nilai float64) (model.IndexPenilaian, error)
	GetIndexPenilaianById(id int) (model.IndexPenilaian, error)
	CreateIndexPenilaian(indexPenilaian model.IndexPenilaian) error
	UpdateIndexPenilaian(indexPenilaian model.IndexPenilaian) error
	DeleteIndexPenilaian(id int) error
}

type indexPenilaianRepository struct {
	dbPenilaian *gorm.DB
}

func NewIndexPenilaianRepo(dbPenilaian *gorm.DB) IndexPenilaianRepository {
	return &indexPenilaianRepository{dbPenilaian}
}

func (r *indexPenilaianRepository) GetIndexPenilaian() ([]model.IndexPenilaian, error) {
	var indexPenilaian []model.IndexPenilaian
	err := r.dbPenilaian.Find(&indexPenilaian).Error
	if err != nil {
		return nil, err
	}
	return indexPenilaian, nil
}
func (r *indexPenilaianRepository) GetIndexPenilaianByNilai(nilai float64) (model.IndexPenilaian, error) {
	var indexPenilaian model.IndexPenilaian
	err := r.dbPenilaian.Where("batas_awal <= ? AND batas_akhir >= ?", nilai, nilai).First(&indexPenilaian).Error
	if err != nil {
		return model.IndexPenilaian{}, err
	}
	return indexPenilaian, nil
}

func (r *indexPenilaianRepository) GetIndexPenilaianById(id int) (model.IndexPenilaian, error) {
	var indexPenilaian model.IndexPenilaian
	err := r.dbPenilaian.First(&indexPenilaian, id).Error
	if err != nil {
		return model.IndexPenilaian{}, err
	}
	return indexPenilaian, nil
}

func (r *indexPenilaianRepository) CreateIndexPenilaian(indexPenilaian model.IndexPenilaian) error {
	err := r.dbPenilaian.Create(&indexPenilaian).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *indexPenilaianRepository) UpdateIndexPenilaian(indexPenilaian model.IndexPenilaian) error {
	err := r.dbPenilaian.Save(&indexPenilaian).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *indexPenilaianRepository) DeleteIndexPenilaian(id int) error {
	err := r.dbPenilaian.Delete(&model.IndexPenilaian{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
