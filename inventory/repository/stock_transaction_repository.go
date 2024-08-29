package repository

import "github.com/iniakunhuda/logistik-tani/inventory/model"

type StockTransactionRepository interface {
	Save(stock model.StockTransaction) error
	Update(stock model.StockTransaction) error
	Delete(stockId int) error
	FindById(stockId int) (*model.StockTransaction, error)
	FindAll() (stocks []model.StockTransaction, err error)
	GetAllByQuery(stock model.StockTransaction) (stocks []model.StockTransaction, err error)
	GetOneByQuery(stock model.StockTransaction) (stockData model.StockTransaction, err error)
}
