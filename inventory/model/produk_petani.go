package model

import (
	"errors"

	"gorm.io/gorm"
)

type ProdukPetaniModel struct {
	DB *gorm.DB
}

// All method will be used to get all records from produks table
func (m *ProdukPetaniModel) All() ([]Produk, error) {
	var produks []Produk
	result := m.DB.Find(&produks)
	return produks, result.Error
}

func (m *ProdukPetaniModel) FindByID(id uint) (*Produk, error) {
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

func (m *ProdukPetaniModel) Insert(produk Produk) (*Produk, error) {
	result := m.DB.Create(&produk)
	if result.Error != nil {
		return nil, result.Error
	}
	return &produk, nil
}

func (m *ProdukPetaniModel) Delete(id uint) error {
	result := m.DB.Delete(&Produk{}, id)
	return result.Error
}

func (m *ProdukPetaniModel) Update(id uint, produk Produk) (*Produk, error) {
	result := m.DB.Model(&Produk{}).Where("id = ?", id).Updates(produk)
	if result.Error != nil {
		return nil, result.Error
	}
	return &produk, nil
}
