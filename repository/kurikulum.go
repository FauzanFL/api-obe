package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type KurikulumRepository interface {
	GetKurikulum() ([]model.Kurikulum, error)
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
