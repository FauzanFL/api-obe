package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type KurikulumRepository interface {
	GetKurikulum() ([]model.Kurikulum, error)
	CreateKurikulum(kurkulum model.Kurikulum) error
	DeleteKurikulum(id int) error
}

type kurikulumRepository struct {
	dbKurikulum *gorm.DB
}

func NewKurikulumRepo(dbKurikulum *gorm.DB) KurikulumRepository {
	return &kurikulumRepository{dbKurikulum}
}

func (k *kurikulumRepository) GetKurikulum() ([]model.Kurikulum, error) {
	var kurikulum []model.Kurikulum
	err := k.dbKurikulum.Find(&kurikulum).Error
	if err != nil {
		return []model.Kurikulum{}, err
	}
	return kurikulum, nil
}

func (k *kurikulumRepository) CreateKurikulum(kurikulum model.Kurikulum) error {
	err := k.dbKurikulum.Create(&kurikulum).Error
	if err != nil {
		return err
	}
	return nil
}

func (k *kurikulumRepository) DeleteKurikulum(id int) error {
	err := k.dbKurikulum.Delete(&model.Kurikulum{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
