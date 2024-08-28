package model

import (
	"errors"

	"gorm.io/gorm"
)

type ProductModel struct {
	DB *gorm.DB
}

// All method will be used to get all records from produks table
func (m *ProductModel) All() ([]Product, error) {
	var produks []Product
	result := m.DB.Find(&produks)
	return produks, result.Error
}

func (m *ProductModel) FindByID(id uint) (*Product, error) {
	var produk Product
	result := m.DB.First(&produk, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &produk, nil
}

func (m *ProductModel) Insert(produk Product) (*Product, error) {
	result := m.DB.Create(&produk)
	if result.Error != nil {
		return nil, result.Error
	}
	return &produk, nil
}

func (m *ProductModel) Delete(id uint) error {
	result := m.DB.Delete(&Product{}, id)
	return result.Error
}

func (m *ProductModel) Update(id uint, produk Product) (*Product, error) {
	result := m.DB.Model(&Product{}).Where("id = ?", id).Updates(produk)
	if result.Error != nil {
		return nil, result.Error
	}
	return &produk, nil
}
