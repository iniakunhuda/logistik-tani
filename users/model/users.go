package model

import (
	"errors"

	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

// All method will be used to get all records from users table
func (m *UserModel) All() ([]User, error) {
	var users []User
	result := m.DB.Find(&users)
	return users, result.Error
}

func (m *UserModel) FindByID(id uint) (*User, error) {
	var user User
	result := m.DB.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (m *UserModel) Insert(user User) (*User, error) {
	result := m.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (m *UserModel) Delete(id uint) error {
	result := m.DB.Delete(&User{}, id)
	return result.Error
}

func (m *UserModel) Update(id uint, user User) (*User, error) {
	result := m.DB.Model(&User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
