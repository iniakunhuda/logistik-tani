package model

import (
	"errors"

	"gorm.io/gorm"
)

type SalesModel struct {
	DB *gorm.DB
}

func (m *SalesModel) All() ([]Sales, error) {
	var sales []Sales
	result := m.DB.Find(&sales)
	return sales, result.Error
}

func (m *SalesModel) FindByID(id uint) (*Sales, error) {
	var sale Sales
	result := m.DB.First(&sale, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &sale, nil
}

func (m *SalesModel) Insert(sale Sales) (*Sales, error) {
	result := m.DB.Create(&sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}

func (m *SalesModel) Delete(id uint) error {
	result := m.DB.Delete(&Sales{}, id)
	return result.Error
}

func (m *SalesModel) Update(id uint, sale Sales) (*Sales, error) {
	result := m.DB.Model(&Sales{}).Where("id = ?", id).Updates(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}
