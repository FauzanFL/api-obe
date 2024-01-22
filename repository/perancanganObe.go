package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PerancanganObeRepository interface {
	GetPerancanganObe() ([]model.PerancanganObeKurikulum, error)
	GetActivePerancanganObe() (model.PerancanganObe, error)
	GetPerancanganObeById(id int) (model.PerancanganObe, error)
	CreatePerancanganObe(perancanganObe model.PerancanganObe) error
	UpdateStatusPerancangan(id int, status string) error
	DiactivatePerancangan() error
	UpdatePerancanganObe(perancanganObe model.PerancanganObe) error
	DeletePerancanganObe(id int) error
}

type perancanganObeRepository struct {
	dbKurikulum *gorm.DB
}

func NewPerancanganObeRepo(dbKurikulum *gorm.DB) PerancanganObeRepository {
	return &perancanganObeRepository{dbKurikulum}
}

func (p *perancanganObeRepository) GetPerancanganObe() ([]model.PerancanganObeKurikulum, error) {
	var perancanganObe []model.PerancanganObeKurikulum
	err := p.dbKurikulum.Model(&model.PerancanganObe{}).Select("perancangan_obe.id, perancangan_obe.nama, perancangan_obe.status, kurikulum.nama as kurikulum").Joins("inner join kurikulum on kurikulum.id = perancangan_obe.kurikulum_id").Scan(&perancanganObe).Error
	if err != nil {
		return []model.PerancanganObeKurikulum{}, err
	}
	return perancanganObe, nil
}

func (p *perancanganObeRepository) GetActivePerancanganObe() (model.PerancanganObe, error) {
	var perancanganObe model.PerancanganObe
	err := p.dbKurikulum.Where("status = ?", "active").First(&perancanganObe).Error
	if err != nil {
		return model.PerancanganObe{}, err
	}
	return perancanganObe, nil
}

func (p *perancanganObeRepository) GetPerancanganObeById(id int) (model.PerancanganObe, error) {
	var perancanganObe model.PerancanganObe
	err := p.dbKurikulum.First(&perancanganObe, id).Error
	if err != nil {
		return model.PerancanganObe{}, err
	}
	return perancanganObe, nil
}

func (p *perancanganObeRepository) CreatePerancanganObe(perancanganObe model.PerancanganObe) error {
	err := p.dbKurikulum.Create(&perancanganObe).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *perancanganObeRepository) UpdateStatusPerancangan(id int, status string) error {
	err := p.dbKurikulum.Model(&model.PerancanganObe{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *perancanganObeRepository) DiactivatePerancangan() error {
	err := p.dbKurikulum.Model(&model.PerancanganObe{}).Where("status = ?", "active").Update("status", "inactive").Error
	if err != nil {
		return err
	}
	return nil
}

func (p *perancanganObeRepository) UpdatePerancanganObe(perancanganObe model.PerancanganObe) error {
	err := p.dbKurikulum.Save(&perancanganObe).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *perancanganObeRepository) DeletePerancanganObe(id int) error {
	err := p.dbKurikulum.Delete(&model.PerancanganObe{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
