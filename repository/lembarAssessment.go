package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type LembarAssessmentRepository interface {
	GetLembarAssessment() ([]model.LembarAssessment, error)
	GetLembarAssessmentById(id int) (model.LembarAssessment, error)
	GetLembarAssessmentByCloId(cLoId int) ([]model.LembarAssessmentWithJenis, error)
	GetLembarAssessmentByPloId(pLoId int) ([]model.LembarAssessment, error)
	GetLembarAssessmentByCloIdAndKeyword(cLoId int, keyword string) ([]model.LembarAssessmentWithJenis, error)
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

func (l *lembarAssessmentRepository) GetLembarAssessmentByPloId(ploId int) ([]model.LembarAssessment, error) {
	var lembarAssessment []model.LembarAssessment
	err := l.dbKurikulum.Model(&model.LembarAssessment{}).Where("clo_id IN (?)", l.dbKurikulum.Model(&model.CLO{}).Select("id").Where("plo_id = ?", ploId)).Scan(&lembarAssessment).Error
	if err != nil {
		return []model.LembarAssessment{}, err
	}
	return lembarAssessment, nil
}

func (l *lembarAssessmentRepository) GetLembarAssessmentByCloId(cloId int) ([]model.LembarAssessmentWithJenis, error) {
	var lembarAssessment []model.LembarAssessmentWithJenis
	err := l.dbKurikulum.Model(&model.LembarAssessment{}).Select("lembar_assessment.id, lembar_assessment.nama, lembar_assessment.deskripsi, lembar_assessment.bobot, lembar_assessment.jenis_id, lembar_assessment.clo_id, jenis_assessment.nama as jenis ").Joins("inner join jenis_assessment on jenis_assessment.id = lembar_assessment.jenis_id").Where("clo_id = ?", cloId).Scan(&lembarAssessment).Error
	if err != nil {
		return []model.LembarAssessmentWithJenis{}, err
	}
	return lembarAssessment, nil
}

func (l *lembarAssessmentRepository) GetLembarAssessmentByCloIdAndKeyword(cloId int, keyword string) ([]model.LembarAssessmentWithJenis, error) {
	key := "%" + keyword + "%"
	var lembarAssessment []model.LembarAssessmentWithJenis
	err := l.dbKurikulum.Model(&model.LembarAssessment{}).Select("lembar_assessment.id, lembar_assessment.nama, lembar_assessment.deskripsi, lembar_assessment.bobot, lembar_assessment.jenis_id, lembar_assessment.clo_id, jenis_assessment.nama as jenis ").Joins("inner join jenis_assessment on jenis_assessment.id = lembar_assessment.jenis_id").Where("clo_id = ? AND nama like ?", cloId, key).Scan(&lembarAssessment).Error
	if err != nil {
		return []model.LembarAssessmentWithJenis{}, err
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
