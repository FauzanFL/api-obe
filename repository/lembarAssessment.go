package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type LembarAssessmentRepository interface {
	GetLembarAssessment() ([]model.LembarAssessment, error)
	GetLembarAssessmentById(id int) (model.LembarAssessment, error)
	CreateLembarAssessment(lembarAssessment model.LembarAssessment) error
	UpdateLembarAssessment(lembarAssessment model.LembarAssessment) error
	DeleteLembarAssessment(id int) error
}

type lembarAssessmentRepository struct {
	dbKurikulum *gorm.DB
}

func NewLembarAssessmentRepo(dbKurikulum *gorm.DB) LembarAssessmentRepository {
	return &lembarAssessmentRepository{dbKurikulum}
}

func (l *lembarAssessmentRepository) GetLembarAssessment() ([]model.LembarAssessment, error) {
	var lembarAssessment []model.LembarAssessment
	err := l.dbKurikulum.Find(&lembarAssessment).Error
	if err != nil {
		return []model.LembarAssessment{}, err
	}
	return lembarAssessment, nil
}

func (l *lembarAssessmentRepository) GetLembarAssessmentById(id int) (model.LembarAssessment, error) {
	var lembarAssessment model.LembarAssessment
	err := l.dbKurikulum.First(&lembarAssessment, id).Error
	if err != nil {
		return model.LembarAssessment{}, err
	}
	return lembarAssessment, nil
}

func (l *lembarAssessmentRepository) CreateLembarAssessment(lembarAssessment model.LembarAssessment) error {
	err := l.dbKurikulum.Create(&lembarAssessment).Error
	if err != nil {
		return err
	}
	return nil
}

func (l *lembarAssessmentRepository) UpdateLembarAssessment(lembarAssessment model.LembarAssessment) error {
	err := l.dbKurikulum.Save(&lembarAssessment).Error
	if err != nil {
		return err
	}
	return nil
}

func (l *lembarAssessmentRepository) DeleteLembarAssessment(id int) error {
	err := l.dbKurikulum.Delete(&model.LembarAssessment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
