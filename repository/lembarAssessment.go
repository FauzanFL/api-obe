package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type LembarAssessmentRepository interface {
	GetLembarAssessment() ([]model.LembarAssessment, error)
}

type lembarAssessmentRepository struct {
	db *gorm.DB
}

func NewLembarAssessmentRepo(db *gorm.DB) LembarAssessmentRepository {
	return &lembarAssessmentRepository{db}
}

func (l *lembarAssessmentRepository) GetLembarAssessment() ([]model.LembarAssessment, error) {
	var lembarAssessment []model.LembarAssessment
	err := l.db.Find(&lembarAssessment).Error
	if err != nil {
		return []model.LembarAssessment{}, err
	}
	return lembarAssessment, nil
}
