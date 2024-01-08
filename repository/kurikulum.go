package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type KurikulumRepository interface {
	GetKurikulum() ([]model.Kurikulum, error)
}

type kurikulumRepository struct {
	db *gorm.DB
}

func NewKurikulumRepo(db *gorm.DB) KurikulumRepository {
	return &kurikulumRepository{db}
}

func (k *kurikulumRepository) GetKurikulum() ([]model.Kurikulum, error) {
	var kurikulum []model.Kurikulum
	err := k.db.Find(&kurikulum).Error
	if err != nil {
		return []model.Kurikulum{}, err
	}
	return kurikulum, nil
}
