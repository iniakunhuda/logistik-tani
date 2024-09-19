package purchaseigmmodel

import (
	"errors"

	"gorm.io/gorm"
)

type PurchaseIgmModel struct {
	DB *gorm.DB
}

func (m *PurchaseIgmModel) All() ([]PurchaseIgm, error) {
	var listData []PurchaseIgm
	result := m.DB.Find(&listData)
	return listData, result.Error
}

func (m *PurchaseIgmModel) FindByID(id uint) (*PurchaseIgm, error) {
	var purchase PurchaseIgm
	result := m.DB.First(&purchase, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &purchase, nil
}

func (m *PurchaseIgmModel) Insert(purchase PurchaseIgm) (*PurchaseIgm, error) {
	result := m.DB.Create(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return &purchase, nil
}

func (m *PurchaseIgmModel) Delete(id uint) error {
	result := m.DB.Delete(&PurchaseIgm{}, id)
	return result.Error
}

func (m *PurchaseIgmModel) Update(id uint, sale PurchaseIgm) (*PurchaseIgm, error) {
	result := m.DB.Model(&PurchaseIgm{}).Where("id = ?", id).Updates(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}
