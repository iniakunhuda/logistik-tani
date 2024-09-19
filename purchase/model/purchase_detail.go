package model

import (
	"errors"

	"gorm.io/gorm"
)

type PurchaseDetailModel struct {
	DB *gorm.DB
}

func (m *PurchaseDetailModel) All() ([]PurchaseDetail, error) {
	var listData []PurchaseDetail
	result := m.DB.Find(&listData)
	return listData, result.Error
}

func (m *PurchaseDetailModel) FindByID(id uint) (*PurchaseDetail, error) {
	var purchase PurchaseDetail
	result := m.DB.First(&purchase, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &purchase, nil
}

func (m *PurchaseDetailModel) Insert(purchase PurchaseDetail) (*PurchaseDetail, error) {
	result := m.DB.Create(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return &purchase, nil
}

func (m *PurchaseDetailModel) Delete(id uint) error {
	result := m.DB.Delete(&PurchaseDetail{}, id)
	return result.Error
}

func (m *PurchaseDetailModel) Update(id uint, sale PurchaseDetail) (*PurchaseDetail, error) {
	result := m.DB.Model(&PurchaseDetail{}).Where("id = ?", id).Updates(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}
