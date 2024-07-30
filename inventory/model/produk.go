package model

import (
	"errors"

	"gorm.io/gorm"
)

type ProdukModel struct {
	DB *gorm.DB
}

// All method will be used to get all records from produks table
func (m *ProdukModel) All() ([]Produk, error) {
	var produks []Produk
	result := m.DB.Find(&produks)
	return produks, result.Error
}

func (m *ProdukModel) FindByID(id uint) (*Produk, error) {
	var produk Produk
	result := m.DB.First(&produk, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &produk, nil
}

func (m *ProdukModel) Insert(produk Produk) (*Produk, error) {
	result := m.DB.Create(&produk)
	if result.Error != nil {
		return nil, result.Error
	}
	return &produk, nil
}

func (m *ProdukModel) Delete(id uint) error {
	result := m.DB.Delete(&Produk{}, id)
	return result.Error
}

func (m *ProdukModel) Update(id uint, produk Produk) (*Produk, error) {
	result := m.DB.Model(&Produk{}).Where("id = ?", id).Updates(produk)
	if result.Error != nil {
		return nil, result.Error
	}
	return &produk, nil
}
