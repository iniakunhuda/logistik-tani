package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/users/model"
	// "github.com/iniakunhuda/logistik-tani/users/request"
	"gorm.io/gorm"
)

type UserLandRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserLandRepositoryImpl(Db *gorm.DB) UserLandRepository {
	return &UserLandRepositoryImpl{Db: Db}
}

func (t *UserLandRepositoryImpl) Delete(userId int) error {
	var user model.UserLand
	result := t.Db.Where("id = ?", userId).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *UserLandRepositoryImpl) FindAll() (users []model.UserLand, err error) {
	var userList []model.UserLand
	result := t.Db.Preload("User").Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	return userList, nil
}

func (t *UserLandRepositoryImpl) FindById(userId int) (*model.UserLand, error) {
	var userResult model.UserLand
	result := t.Db.Preload("User").Find(&userResult, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("land is not found")
	}
	return &userResult, nil
}

func (t *UserLandRepositoryImpl) Save(user model.UserLand) error {
	result := t.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *UserLandRepositoryImpl) Update(user model.UserLand) error {
	result := t.Db.Model(&user).Updates(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *UserLandRepositoryImpl) GetAllByQuery(user model.UserLand) (users []model.UserLand, err error) {
	var userList []model.UserLand
	result := t.Db.Where(&user).Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	return userList, nil
}

func (t *UserLandRepositoryImpl) GetOneByQuery(user model.UserLand) (userData model.UserLand, err error) {
	result := t.Db.Where(&user).First(&userData)
	if result.Error != nil {
		return model.UserLand{}, result.Error
	}
	return userData, nil
}
