package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type TahunAjaranRepository interface {
	GetTahunAjaran() ([]model.TahunAjaran, error)
	GetTahunAjaranById(id int) (model.TahunAjaran, error)
	GetTahunAjaranByKeyword(keyword string) ([]model.TahunAjaran, error)
	GetTahunAjaranByMonth(currentMonth int) ([]model.TahunAjaran, error)
	CreateTahunAjaran(tahunAjar model.TahunAjaran) error
	UpdateTahunAjaran(tahunAjar model.TahunAjaran) error
	DeleteTahunAjaran(id int) error
}

type tahunAjaranRepository struct {
	dbPenilaian *gorm.DB
}

func NewTahunAjaranRepo(dbPenilaian *gorm.DB) TahunAjaranRepository {
	return &tahunAjaranRepository{dbPenilaian}
}

func (p *tahunAjaranRepository) GetTahunAjaran() ([]model.TahunAjaran, error) {
	var tahunAjaran []model.TahunAjaran
	err := p.dbPenilaian.Find(&tahunAjaran).Error
	if err != nil {
		return []model.TahunAjaran{}, err
	}
	return tahunAjaran, nil
}

func (p *tahunAjaranRepository) GetTahunAjaranById(id int) (model.TahunAjaran, error) {
	var tahunAjaran model.TahunAjaran
	err := p.dbPenilaian.First(&tahunAjaran, id).Error
	if err != nil {
		return model.TahunAjaran{}, err
	}
	return tahunAjaran, nil
}

func (p *tahunAjaranRepository) GetTahunAjaranByKeyword(keyword string) ([]model.TahunAjaran, error) {
	var tahunAjaran []model.TahunAjaran
	err := p.dbPenilaian.Where("tahun LIKE ? OR semester LIKE ?", "%"+keyword+"%", "%"+keyword+"%").Find(&tahunAjaran).Error
	if err != nil {
		return []model.TahunAjaran{}, err
	}
	return tahunAjaran, nil
}

func (p *tahunAjaranRepository) GetTahunAjaranByMonth(currentMonth int) ([]model.TahunAjaran, error) {
	var tahunAjaran []model.TahunAjaran
	err := p.dbPenilaian.Where("bulan_mulai <= ? AND bulan_selesai >= ?", currentMonth, currentMonth).Find(&tahunAjaran).Error
	if err != nil {
		return []model.TahunAjaran{}, err
	}
	return tahunAjaran, nil
}

func (p *tahunAjaranRepository) CreateTahunAjaran(tahunAjar model.TahunAjaran) error {
	err := p.dbPenilaian.Create(&tahunAjar).Error
	return err
}

func (p *tahunAjaranRepository) UpdateTahunAjaran(tahunAjar model.TahunAjaran) error {
	err := p.dbPenilaian.Save(&tahunAjar).Error
	return err
}

func (p *tahunAjaranRepository) DeleteTahunAjaran(id int) error {
	err := p.dbPenilaian.Delete(&model.TahunAjaran{}, id).Error
	return err
}
