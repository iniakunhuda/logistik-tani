package model

import (
	"errors"

	"gorm.io/gorm"
)

type UserLandModel struct {
	DB *gorm.DB
}

// All method will be used to get all records from users table
func (m *UserLandModel) All() ([]UserLand, error) {
	var users []UserLand
	result := m.DB.Find(&users)
	return users, result.Error
}

func (m *UserLandModel) FindByID(id uint) (*UserLand, error) {
	var user UserLand
	result := m.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (m *UserLandModel) Insert(user UserLand) (*UserLand, error) {
	result := m.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (m *UserLandModel) Delete(id uint) error {
	result := m.DB.Delete(&UserLand{}, id)
	return result.Error
}

func (m *UserLandModel) Update(id uint, user UserLand) (*UserLand, error) {
	result := m.DB.Model(&UserLand{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
