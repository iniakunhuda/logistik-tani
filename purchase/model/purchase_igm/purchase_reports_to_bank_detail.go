package purchaseigm

import (
	"errors"

	"gorm.io/gorm"
)

type PurchaseReportsToBankDetailModel struct {
	DB *gorm.DB
}

func (m *PurchaseReportsToBankDetailModel) All() ([]PurchaseReportsToBankDetail, error) {
	var listData []PurchaseReportsToBankDetail
	result := m.DB.Find(&listData)
	return listData, result.Error
}

func (m *PurchaseReportsToBankDetailModel) FindByID(id uint) (*PurchaseReportsToBankDetail, error) {
	var purchase PurchaseReportsToBankDetail
	result := m.DB.First(&purchase, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &purchase, nil
}

func (m *PurchaseReportsToBankDetailModel) Insert(purchase PurchaseReportsToBankDetail) (*PurchaseReportsToBankDetail, error) {
	result := m.DB.Create(&purchase)
	if result.Error != nil {
		return nil, result.Error
	}
	return &purchase, nil
}

func (m *PurchaseReportsToBankDetailModel) Delete(id uint) error {
	result := m.DB.Delete(&PurchaseReportsToBankDetail{}, id)
	return result.Error
}

func (m *PurchaseReportsToBankDetailModel) Update(id uint, sale PurchaseReportsToBankDetail) (*PurchaseReportsToBankDetail, error) {
	result := m.DB.Model(&PurchaseReportsToBankDetail{}).Where("id = ?", id).Updates(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}
