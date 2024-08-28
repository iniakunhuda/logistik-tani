package model

import (
	"errors"

	"gorm.io/gorm"
)

type StockTransactionModel struct {
	DB *gorm.DB
}

// All method will be used to get all records from listData table
func (m *StockTransactionModel) All() ([]StockTransaction, error) {
	var listData []StockTransaction
	result := m.DB.Find(&listData)
	return listData, result.Error
}

func (m *StockTransactionModel) FindByID(id uint) (*StockTransaction, error) {
	var data StockTransaction
	result := m.DB.First(&data, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, result.Error
	}
	return &data, nil
}

func (m *StockTransactionModel) Insert(data StockTransaction) (*StockTransaction, error) {
	result := m.DB.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (m *StockTransactionModel) Delete(id uint) error {
	result := m.DB.Delete(&StockTransaction{}, id)
	return result.Error
}

func (m *StockTransactionModel) Update(id uint, data StockTransaction) (*StockTransaction, error) {
	result := m.DB.Model(&StockTransaction{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}
