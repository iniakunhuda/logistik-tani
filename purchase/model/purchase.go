package model

import (
	"errors"

	"gorm.io/gorm"
)

type PurchaseModel struct {
	DB *gorm.DB
}

func (m *PurchaseModel) All() ([]Purchase, error) {
	var sales []Purchase
	result := m.DB.Find(&sales)
	return sales, result.Error
}

func (m *PurchaseModel) FindByID(id uint) (*Purchase, error) {
	var sale Purchase
	result := m.DB.First(&sale, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &sale, nil
}

func (m *PurchaseModel) Insert(sale Purchase) (*Purchase, error) {
	result := m.DB.Create(&sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}

func (m *PurchaseModel) Delete(id uint) error {
	result := m.DB.Delete(&Purchase{}, id)
	return result.Error
}

func (m *PurchaseModel) Update(id uint, sale Purchase) (*Purchase, error) {
	result := m.DB.Model(&Purchase{}).Where("id = ?", id).Updates(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}
