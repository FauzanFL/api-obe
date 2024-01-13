package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	GetMahasiswa() ([]model.Mahasiswa, error)
	GetMahasiswaByMataKuliah(mkId int) ([]model.Mahasiswa, error)
	GetMahasiswaByKelasMataKuliah(mkId int, kelasId int) ([]model.Mahasiswa, error)
}

type mahasiswaRepository struct {
	dbPenilaian *gorm.DB
}

func NewMahasiswaRepo(dbPenilaian *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepository{dbPenilaian}
}

func (m *mahasiswaRepository) GetMahasiswa() ([]model.Mahasiswa, error) {
	var mahasiswa []model.Mahasiswa
	err := m.dbPenilaian.Find(&mahasiswa).Error
	if err != nil {
		return []model.Mahasiswa{}, err
	}
	return mahasiswa, nil
}

func (m *mahasiswaRepository) GetMahasiswaByMataKuliah(mkId int) ([]model.Mahasiswa, error) {
	var mahasiswa []model.Mahasiswa
	err := m.dbPenilaian.Model(&model.Mahasiswa{}).Where("id IN ?", m.dbPenilaian.Table("mk_mahasiswa").Where("mk_id = ?", mkId).Select("mhs_id")).Error
	if err != nil {
		return []model.Mahasiswa{}, err
	}
	return mahasiswa, nil
}

func (m *mahasiswaRepository) GetMahasiswaByKelasMataKuliah(mkId int, kelasId int) ([]model.Mahasiswa, error) {
	var mahasiswa []model.Mahasiswa
	err := m.dbPenilaian.Model(&model.Mahasiswa{}).Where("id IN ?", m.dbPenilaian.Table("mk_mahasiswa").Where("mk_id = ? AND kelas_id = ?", mkId, kelasId).Select("mhs_id"), kelasId).Error
	if err != nil {
		return []model.Mahasiswa{}, err
	}
	return mahasiswa, nil
}
