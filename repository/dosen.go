package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type DosenRepository interface {
	GetDosen() ([]model.Dosen, error)
	GetDosenById(id int) (model.Dosen, error)
	Add(dosen model.Dosen) error
	Delete(id int) error
	UpdateByUser(dosen model.Dosen, userId int) error
}

type dosenRepository struct {
	dbUser *gorm.DB
}

func NewDosenRepo(dbUser *gorm.DB) DosenRepository {
	return &dosenRepository{dbUser}
}

func (d *dosenRepository) GetDosen() ([]model.Dosen, error) {
	var dosen []model.Dosen
	err := d.dbUser.Find(&dosen).Error
	if err != nil {
		return []model.Dosen{}, err
	}
	return dosen, nil
}

func (d *dosenRepository) GetDosenById(id int) (model.Dosen, error) {
	var dosen model.Dosen
	err := d.dbUser.Where("id =?", id).First(&dosen).Error
	if err != nil {
		return model.Dosen{}, err
	}
	return dosen, nil
}

func (d *dosenRepository) Add(dosen model.Dosen) error {
	err := d.dbUser.Create(&dosen).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *dosenRepository) Delete(id int) error {
	err := d.dbUser.Where("id =?", id).Delete(&model.Dosen{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *dosenRepository) UpdateByUser(dosen model.Dosen, userId int) error {
	err := d.dbUser.Model(&model.Dosen{}).Where("user_id =?", userId).Updates(&dosen).Error
	if err != nil {
		return err
	}
	return nil
}
