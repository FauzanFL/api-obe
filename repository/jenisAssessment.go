package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type JenisAssessmentRepository interface {
	GetJenisAssessment() ([]model.JenisAssessment, error)
}

type jenisAssessmentRepository struct {
	db *gorm.DB
}

func NewJenisAssessmentRepo(db *gorm.DB) JenisAssessmentRepository {
	return &jenisAssessmentRepository{db}
}

func (j *jenisAssessmentRepository) GetJenisAssessment() ([]model.JenisAssessment, error) {
	var jenisAssessment []model.JenisAssessment
	err := j.db.Find(&jenisAssessment).Error
	if err != nil {
		return []model.JenisAssessment{}, err
	}
	return jenisAssessment, nil
}
