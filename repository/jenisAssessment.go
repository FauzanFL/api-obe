package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type JenisAssessmentRepository interface {
	GetJenisAssessment() ([]model.JenisAssessment, error)
	CreateJenisAssessment(jenisAssessment model.JenisAssessment) error
	UpdateJenisAssessment(jenisAssessment model.JenisAssessment) error
	DeleteJenisAssessment(id int) error
}

type jenisAssessmentRepository struct {
	dbKurikulum *gorm.DB
}

func NewJenisAssessmentRepo(dbKurikulum *gorm.DB) JenisAssessmentRepository {
	return &jenisAssessmentRepository{dbKurikulum}
}

func (j *jenisAssessmentRepository) GetJenisAssessment() ([]model.JenisAssessment, error) {
	var jenisAssessment []model.JenisAssessment
	err := j.dbKurikulum.Find(&jenisAssessment).Error
	if err != nil {
		return []model.JenisAssessment{}, err
	}
	return jenisAssessment, nil
}

func (j *jenisAssessmentRepository) CreateJenisAssessment(jenisAssessment model.JenisAssessment) error {
	err := j.dbKurikulum.Create(&jenisAssessment).Error
	if err != nil {
		return err
	}
	return nil
}

func (j *jenisAssessmentRepository) UpdateJenisAssessment(jenisAssessment model.JenisAssessment) error {
	err := j.dbKurikulum.Save(&jenisAssessment).Error
	if err != nil {
		return err
	}
	return nil
}

func (j *jenisAssessmentRepository) DeleteJenisAssessment(id int) error {
	err := j.dbKurikulum.Delete(&model.JenisAssessment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
