package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/users/model"
	// "github.com/iniakunhuda/logistik-tani/users/request"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

func (t *UserRepositoryImpl) Delete(userId int) error {
	var user model.User
	result := t.Db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *UserRepositoryImpl) FindAll() (users []model.User, err error) {
	var userList []model.User
	result := t.Db.Preload("Lands").Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	return userList, nil
}

func (t *UserRepositoryImpl) FindById(userId int) (*model.User, error) {
	var userResult model.User
	result := t.Db.Preload("Lands").Find(&userResult, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("user is not found")
	}
	return &userResult, nil
}

func (t *UserRepositoryImpl) Save(user model.User) error {
	result := t.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *UserRepositoryImpl) Update(user model.User) error {
	result := t.Db.Model(&user).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *UserRepositoryImpl) GetAllByQuery(user model.User) (users []model.User, err error) {
	var userList []model.User
	result := t.Db.Where(&user).Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	return userList, nil
}

func (t *UserRepositoryImpl) GetOneByQuery(user model.User) (userData model.User, err error) {
	result := t.Db.Where(&user).First(&userData)
	if result.Error != nil {
		return model.User{}, result.Error
	}
	return userData, nil
}
