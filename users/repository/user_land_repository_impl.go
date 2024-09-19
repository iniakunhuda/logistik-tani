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

func (t *UserLandRepositoryImpl) Delete(landId int) error {
	var land model.UserLand
	result := t.Db.Where("id = ?", landId).Delete(&land)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *UserLandRepositoryImpl) FindAll() (users []model.UserLand, err error) {
	var landList []model.UserLand
	result := t.Db.Preload("User").Find(&landList)
	if result.Error != nil {
		return nil, result.Error
	}
	return landList, nil
}

func (t *UserLandRepositoryImpl) FindById(landId int) (*model.UserLand, error) {
	var landResult model.UserLand
	result := t.Db.Preload("User").Find(&landResult, landId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("land is not found")
	}
	return &landResult, nil
}

func (t *UserLandRepositoryImpl) Save(land model.UserLand) error {
	result := t.Db.Create(&land)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *UserLandRepositoryImpl) Update(land model.UserLand) error {
	result := t.Db.Model(&land).Updates(land)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *UserLandRepositoryImpl) GetAllByQuery(land model.UserLand) (users []model.UserLand, err error) {
	var landList []model.UserLand
	result := t.Db.Where(&land).Preload("User").Find(&landList)
	if result.Error != nil {
		return nil, result.Error
	}
	return landList, nil
}

func (t *UserLandRepositoryImpl) GetOneByQuery(land model.UserLand) (userData model.UserLand, err error) {
	result := t.Db.Where(&land).Preload("User").First(&userData)
	if result.Error != nil {
		return model.UserLand{}, result.Error
	}
	return userData, nil
}
