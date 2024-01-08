package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type CloRepository interface {
	GetClo() ([]model.CLO, error)
	GetCloById(id int) (model.CLO, error)
	CreateClo(clo model.CLO) error
	UpdateClo(clo model.CLO) error
	DeleteClo(id int) error
}

type cloRepository struct {
	dbKurikulum *gorm.DB
}

func NewCloRepo(dbKurikulum *gorm.DB) CloRepository {
	return &cloRepository{dbKurikulum}
}

func (c *cloRepository) GetClo() ([]model.CLO, error) {
	var clo []model.CLO
	err := c.dbKurikulum.Find(&clo).Error
	if err != nil {
		return []model.CLO{}, err
	}
	return clo, nil
}

func (c *cloRepository) GetCloById(id int) (model.CLO, error) {
	var clo model.CLO
	err := c.dbKurikulum.First(&clo, id).Error
	if err != nil {
		return model.CLO{}, err
	}
	return clo, nil
}

func (c *cloRepository) CreateClo(clo model.CLO) error {
	err := c.dbKurikulum.Create(&clo).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cloRepository) UpdateClo(clo model.CLO) error {
	err := c.dbKurikulum.Save(&clo).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *cloRepository) DeleteClo(id int) error {
	err := c.dbKurikulum.Delete(&model.CLO{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
