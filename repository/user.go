package repository

import (
	"api-obe/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser() ([]model.User, error)
	GetUserById(id int) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
	Add(user model.User) error
	Update(user model.User) error
	Delete(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) GetUser() ([]model.User, error) {
	var users []model.User
	err := u.db.Find(&users).Error
	if err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func (u *userRepository) GetUserById(id int) (model.User, error) {
	var user model.User
	err := u.db.First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return user, err
}

func (u *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := u.db.Find(&user, "email = ?", email).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.User{}, nil
		}
		return model.User{}, err
	}
	return user, err
}

func (u *userRepository) Add(user model.User) error {
	return u.db.Create(&user).Error
}

func (u *userRepository) Update(user model.User) error {
	return u.db.Save(&user).Error
}

func (u *userRepository) Delete(id int) error {
	return u.db.Where("id = ?", id).Delete(&model.User{}).Error
}
