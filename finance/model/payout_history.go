package model

import (
	"errors"

	"gorm.io/gorm"
)

type PayoutHistoryModel struct {
	DB *gorm.DB
}

func (m *PayoutHistoryModel) All() ([]PayoutHistory, error) {
	var sales []PayoutHistory
	result := m.DB.Find(&sales)
	return sales, result.Error
}

func (m *PayoutHistoryModel) FindByID(id uint) (*PayoutHistory, error) {
	var sale PayoutHistory
	result := m.DB.First(&sale, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &sale, nil
}

func (m *PayoutHistoryModel) Insert(sale PayoutHistory) (*PayoutHistory, error) {
	result := m.DB.Create(&sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}

func (m *PayoutHistoryModel) Delete(id uint) error {
	result := m.DB.Delete(&PayoutHistory{}, id)
	return result.Error
}

func (m *PayoutHistoryModel) Update(id uint, sale PayoutHistory) (*PayoutHistory, error) {
	result := m.DB.Model(&PayoutHistory{}).Where("id = ?", id).Updates(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}
