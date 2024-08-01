package model

import (
	"errors"

	"gorm.io/gorm"
)

type SalesIgmDetailModel struct {
	DB *gorm.DB
}

func (m *SalesIgmDetailModel) All() ([]SalesIgmDetail, error) {
	var sales []SalesIgmDetail
	result := m.DB.Find(&sales)
	return sales, result.Error
}

func (m *SalesIgmDetailModel) FindByID(id uint) (*SalesIgmDetail, error) {
	var sale SalesIgmDetail
	result := m.DB.First(&sale, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &sale, nil
}

func (m *SalesIgmDetailModel) Insert(sale SalesIgmDetail) (*SalesIgmDetail, error) {
	result := m.DB.Create(&sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}

func (m *SalesIgmDetailModel) Delete(id uint) error {
	result := m.DB.Delete(&SalesIgmDetail{}, id)
	return result.Error
}

func (m *SalesIgmDetailModel) Update(id uint, sale SalesIgmDetail) (*SalesIgmDetail, error) {
	result := m.DB.Model(&SalesIgmDetail{}).Where("id = ?", id).Updates(sale)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sale, nil
}
