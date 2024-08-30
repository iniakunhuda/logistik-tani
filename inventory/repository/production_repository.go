package repository

import "github.com/iniakunhuda/logistik-tani/inventory/model"

type ProductionRepository interface {
	Save(stock model.Production) error
	Update(stock model.Production) error
	Delete(stockId int) error
	FindById(stockId int) (*model.Production, error)
	FindAll() (stocks []model.Production, err error)
	GetAllByQuery(stock model.Production) (stocks []model.Production, err error)
	GetOneByQuery(stock model.Production) (stockData model.Production, err error)
}
