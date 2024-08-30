package purchaseigm

import (
	"errors"

	"gorm.io/gorm"
)

type PurchaseReportsToBankModel struct {
	DB *gorm.DB
}

func (m *PurchaseReportsToBankModel) All() ([]PurchaseReportsToBank, error) {
	var listData []PurchaseReportsToBank
	result := m.DB.Find(&listData)
	return listData, result.Error
}

func (m *PurchaseReportsToBankModel) FindByID(id uint) (*PurchaseReportsToBank, error) {
	var purchase PurchaseReportsToBank
	result := m.DB.First(&purchase, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &purchase, nil
}

func (m *PurchaseReportsToBankModel) Insert(purchase PurchaseReportsToBank) (*PurchaseReportsToBank, error) {
	result := m.DB.Create(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return &purchase, nil
}

func (m *PurchaseReportsToBankModel) Delete(id uint) error {
	result := m.DB.Delete(&PurchaseReportsToBank{}, id)
	return result.Error
}

func (m *PurchaseReportsToBankModel) Update(id uint, sale PurchaseReportsToBank) (*PurchaseReportsToBank, error) {
	result := m.DB.Model(&PurchaseReportsToBank{}).Where("id = ?", id).Updates(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}
