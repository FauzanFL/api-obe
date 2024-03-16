package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type MataKuliahRepository interface {
	GetMataKuliah() ([]model.MataKuliah, error)
	GetMataKuliahById(id int) (model.MataKuliah, error)
	GetCLOFromMataKuliah(id int) ([]model.CLO, error)
	GetPLOFromMataKuliah(id int) ([]model.PLO, error)
	GetMataKuliahByObeIdAndTahunId(obeId int, tahunId int) ([]model.MataKuliah, error)
	GetMataKuliahByObeIdTahunIdAndKeyword(obeId int, tahunId int, keyword string) ([]model.MataKuliah, error)
	GetMataKuliahByDosenObeId(obeId int, dosenId int) ([]model.MataKuliah, error)
	GetMataKuliahByDosenObeIdAndKeyword(obeId int, dosenId int, keyword string) ([]model.MataKuliah, error)
	CreateMataKuliah(mataKuliah model.MataKuliah) error
	UpdateMataKuliah(mataKuliah model.MataKuliah) error
	DeleteMataKuliah(id int) error
}

type mataKuliahRepository struct {
	dbKurikulum *gorm.DB
}

func NewMataKuliahRepo(dbKurikulum *gorm.DB) MataKuliahRepository {
	return &mataKuliahRepository{dbKurikulum}
}

func (m *mataKuliahRepository) GetMataKuliah() ([]model.MataKuliah, error) {
	var mataKuliah []model.MataKuliah
	err := m.dbKurikulum.Find(&mataKuliah).Error
	if err != nil {
		return []model.MataKuliah{}, err
	}
	return mataKuliah, nil
}

func (m *mataKuliahRepository) GetMataKuliahById(id int) (model.MataKuliah, error) {
	var mataKuliah model.MataKuliah
	err := m.dbKurikulum.First(&mataKuliah, id).Error
	if err != nil {
		return model.MataKuliah{}, err
	}
	return mataKuliah, nil
}

func (m *mataKuliahRepository) GetCLOFromMataKuliah(id int) ([]model.CLO, error) {
	var clo []model.CLO
	err := m.dbKurikulum.Model(&model.CLO{}).Where("mk_id =?", id).Find(&clo).Error
	if err != nil {
		return []model.CLO{}, err
	}
	return clo, nil
}

func (m *mataKuliahRepository) GetPLOFromMataKuliah(id int) ([]model.PLO, error) {
	var plo []model.PLO
	err := m.dbKurikulum.Model(&model.PLO{}).Where("id IN (?)", m.dbKurikulum.Table("clo").Where("mk_id = ?", id).Select("plo_id")).Find(&plo).Error
	if err != nil {
		return []model.PLO{}, err
	}
	return plo, nil
}

func (m *mataKuliahRepository) GetMataKuliahByObeIdAndTahunId(obeId int, tahunId int) ([]model.MataKuliah, error) {
	var mataKuliah []model.MataKuliah
	err := m.dbKurikulum.Model(&model.MataKuliah{}).Where("obe_id = ? AND tahun_ajaran_id = ?", obeId, tahunId).Find(&mataKuliah).Error
	if err != nil {
		return []model.MataKuliah{}, err
	}
	return mataKuliah, nil
}

func (m *mataKuliahRepository) GetMataKuliahByObeIdTahunIdAndKeyword(obeId int, tahunId int, keyword string) ([]model.MataKuliah, error) {
	key := "%" + keyword + "%"
	var mataKuliah []model.MataKuliah
	err := m.dbKurikulum.Model(&model.MataKuliah{}).Where("obe_id = ? AND tahun_ajaran_id = ? AND kode_mk like ? OR nama like ?", obeId, tahunId, key, key).Find(&mataKuliah).Error
	if err != nil {
		return []model.MataKuliah{}, err
	}
	return mataKuliah, nil
}

func (m *mataKuliahRepository) GetMataKuliahByDosenObeId(obeId int, dosenId int) ([]model.MataKuliah, error) {
	var mataKuliah []model.MataKuliah
	err := m.dbKurikulum.Model(&model.MataKuliah{}).Where("obe_id = ? AND id IN (?)", obeId, m.dbKurikulum.Table("plotting_dosen_mk").Where("dosen_id = ?", dosenId).Select("mk_id")).Find(&mataKuliah).Error
	if err != nil {
		return []model.MataKuliah{}, err
	}
	return mataKuliah, nil
}

func (m *mataKuliahRepository) GetMataKuliahByDosenObeIdAndKeyword(obeId int, dosenId int, keyword string) ([]model.MataKuliah, error) {
	key := "%" + keyword + "%"
	var mataKuliah []model.MataKuliah
	err := m.dbKurikulum.Model(&model.MataKuliah{}).Where("obe_id = ? AND id IN (?) AND (nama like ? OR kode_mk like ?)", obeId, m.dbKurikulum.Table("plotting_dosen_mk").Where("dosen_id = ?", dosenId).Select("mk_id"), key, key).Find(&mataKuliah).Error
	if err != nil {
		return []model.MataKuliah{}, err
	}
	return mataKuliah, nil
}

func (m *mataKuliahRepository) CreateMataKuliah(mataKuliah model.MataKuliah) error {
	err := m.dbKurikulum.Create(&mataKuliah).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *mataKuliahRepository) UpdateMataKuliah(mataKuliah model.MataKuliah) error {
	err := m.dbKurikulum.Save(&mataKuliah).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *mataKuliahRepository) DeleteMataKuliah(id int) error {
	err := m.dbKurikulum.Delete(&model.MataKuliah{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
