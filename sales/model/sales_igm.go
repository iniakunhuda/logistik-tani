package model

import (
	"errors"

	"gorm.io/gorm"
)

type SalesIgmModel struct {
	DB *gorm.DB
}

func (m *SalesIgmModel) All() ([]SalesIgm, error) {
	var sales []SalesIgm
	result := m.DB.Find(&sales)
	return sales, result.Error
}

func (m *SalesIgmModel) FindByID(id uint) (*SalesIgm, error) {
	var sale SalesIgm
	result := m.DB.First(&sale, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &sale, nil
}

func (m *SalesIgmModel) Insert(sale SalesIgm) (*SalesIgm, error) {
	result := m.DB.Create(&sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}

func (m *SalesIgmModel) Delete(id uint) error {
	result := m.DB.Delete(&SalesIgm{}, id)
	return result.Error
}

func (m *SalesIgmModel) Update(id uint, sale SalesIgm) (*SalesIgm, error) {
	result := m.DB.Model(&SalesIgm{}).Where("id = ?", id).Updates(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}
