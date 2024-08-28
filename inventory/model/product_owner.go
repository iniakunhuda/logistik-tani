package model

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/inventory/response"
	"gorm.io/gorm"
)

type ProductOwnerModel struct {
	DB *gorm.DB
}

// All method will be used to get all records from listData table
func (m *ProductOwnerModel) All() ([]ProductOwner, error) {
	var listData []ProductOwner
	result := m.DB.Find(&listData)
	return listData, result.Error
}

func (m *ProductOwnerModel) FindByID(id uint) (*ProductOwner, error) {
	var data ProductOwner
	result := m.DB.First(&data, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &data, nil
}

func (m *ProductOwnerModel) Insert(data ProductOwner) (*ProductOwner, error) {
	result := m.DB.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (m *ProductOwnerModel) Delete(id uint) error {
	result := m.DB.Delete(&ProductOwner{}, id)
	return result.Error
}

func (m *ProductOwnerModel) Update(id uint, data ProductOwner) (*ProductOwner, error) {
	result := m.DB.Model(&ProductOwner{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (m *ProductOwnerModel) toResponse() response.ProductResponse {
	return response.ProductResponse{}
}
