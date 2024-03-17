package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type PlottingDosenMkRepository interface {
	GetPlottingDosenMk() ([]model.PlottingDosenMk, error)
	GetPlottingDosenMkByObeIdAndTahunId(obeId int, tahunId int) ([]model.PlottingDosenMk, error)
	GetPlottingDosenByMkId(mkId int) ([]model.PlottingDosenMk, error)
	GetPlottingDosenByMkIdAndDosenId(mkId int, dosenId int) ([]model.PlottingDosenMk, error)
	CreatePlottingDosenMk(plottingDosenMk model.PlottingDosenMk) error
	DeletePlottingDosenMk(id int) error
}

type plottingDosenMkRepository struct {
	dbKurikulum *gorm.DB
}

func NewPlottingDosenMkRepo(dbKurikulum *gorm.DB) PlottingDosenMkRepository {
	return &plottingDosenMkRepository{dbKurikulum}
}

func (p *plottingDosenMkRepository) GetPlottingDosenMk() ([]model.PlottingDosenMk, error) {
	var plottingDosenMk []model.PlottingDosenMk
	err := p.dbKurikulum.Find(&plottingDosenMk).Error
	if err != nil {
		return []model.PlottingDosenMk{}, err
	}
	return plottingDosenMk, nil
}

func (p *plottingDosenMkRepository) GetPlottingDosenMkByObeIdAndTahunId(obeId int, tahunId int) ([]model.PlottingDosenMk, error) {
	var plottingDosenMk []model.PlottingDosenMk
	err := p.dbKurikulum.Where("mk_id IN (?)", p.dbKurikulum.Model(&model.MataKuliah{}).Select("id").Where("obe_id = ? AND tahun_ajaran_id = ?", obeId, tahunId)).Find(&plottingDosenMk).Error
	if err != nil {
		return []model.PlottingDosenMk{}, err
	}
	return plottingDosenMk, nil
}

func (p *plottingDosenMkRepository) GetPlottingDosenByMkId(mkId int) ([]model.PlottingDosenMk, error) {
	var plottingDosenMk []model.PlottingDosenMk
	err := p.dbKurikulum.Where("mk_id =?", mkId).Find(&plottingDosenMk).Error
	if err != nil {
		return []model.PlottingDosenMk{}, err
	}
	return plottingDosenMk, nil
}

func (p *plottingDosenMkRepository) GetPlottingDosenByMkIdAndDosenId(mkId int, dosenId int) ([]model.PlottingDosenMk, error) {
	var plottingDosenMk []model.PlottingDosenMk
	err := p.dbKurikulum.Where("mk_id =? AND dosen_id = ?", mkId, dosenId).Find(&plottingDosenMk).Error
	if err != nil {
		return []model.PlottingDosenMk{}, err
	}
	return plottingDosenMk, nil
}

func (p *plottingDosenMkRepository) CreatePlottingDosenMk(plottingDosenMk model.PlottingDosenMk) error {
	err := p.dbKurikulum.Create(&plottingDosenMk).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *plottingDosenMkRepository) DeletePlottingDosenMk(id int) error {
	err := p.dbKurikulum.Delete(&model.PlottingDosenMk{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
