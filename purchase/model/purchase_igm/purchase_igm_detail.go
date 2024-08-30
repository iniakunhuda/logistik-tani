package purchaseigm

import (
	"errors"

	"gorm.io/gorm"
)

type PurchaseIgmDetailModel struct {
	DB *gorm.DB
}

func (m *PurchaseIgmDetailModel) All() ([]PurchaseIgmDetail, error) {
	var listData []PurchaseIgmDetail
	result := m.DB.Find(&listData)
	return listData, result.Error
}

func (m *PurchaseIgmDetailModel) FindByID(id uint) (*PurchaseIgmDetail, error) {
	var purchase PurchaseIgmDetail
	result := m.DB.First(&purchase, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &purchase, nil
}

func (m *PurchaseIgmDetailModel) Insert(purchase PurchaseIgmDetail) (*PurchaseIgmDetail, error) {
	result := m.DB.Create(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return &purchase, nil
}

func (m *PurchaseIgmDetailModel) Delete(id uint) error {
	result := m.DB.Delete(&PurchaseIgmDetail{}, id)
	return result.Error
}

func (m *PurchaseIgmDetailModel) Update(id uint, sale PurchaseIgmDetail) (*PurchaseIgmDetail, error) {
	result := m.DB.Model(&PurchaseIgmDetail{}).Where("id = ?", id).Updates(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}
