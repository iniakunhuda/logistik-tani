package repository

import (
	"errors"

	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"gorm.io/gorm"
)

type StockTransactionRepositoryImpl struct {
	Db *gorm.DB
}

func NewStockTransactionRepositoryImpl(Db *gorm.DB) StockTransactionRepository {
	return &StockTransactionRepositoryImpl{Db: Db}
}

func (t *StockTransactionRepositoryImpl) Save(stock model.StockTransaction) error {
	result := t.Db.Create(&stock)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *StockTransactionRepositoryImpl) FindAll() (stocks []model.StockTransaction, err error) {
	var dataList []model.StockTransaction
	result := t.Db.Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *StockTransactionRepositoryImpl) FindById(stockId int) (*model.StockTransaction, error) {
	var dataRow model.StockTransaction
	result := t.Db.Find(&dataRow, stockId)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("stock not found")
	}
	return &dataRow, nil
}

func (t *StockTransactionRepositoryImpl) Delete(dataId int) error {
	var data model.StockTransaction
	result := t.Db.Where("id = ?", dataId).Delete(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *StockTransactionRepositoryImpl) Update(stock model.StockTransaction) error {
	result := t.Db.Model(&stock).Updates(stock)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *StockTransactionRepositoryImpl) GetAllByQuery(stock model.StockTransaction) (stocks []model.StockTransaction, err error) {
	var dataList []model.StockTransaction
	result := t.Db.Where(&stock).Find(&dataList)
	if result.Error != nil {
		return nil, result.Error
	}
	return dataList, nil
}

func (t *StockTransactionRepositoryImpl) GetOneByQuery(stock model.StockTransaction) (stockData model.StockTransaction, err error) {
	result := t.Db.Where(&stock).First(&stockData)
	if result.Error != nil {
		return model.StockTransaction{}, result.Error
	}
	return stockData, nil
}

func (t *StockTransactionRepositoryImpl) DeleteByQuery(stock model.StockTransaction) error {
	result := t.Db.Where(&stock).Delete(&stock)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
