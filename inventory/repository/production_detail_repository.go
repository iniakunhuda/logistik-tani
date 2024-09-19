package repository

import "github.com/iniakunhuda/logistik-tani/inventory/model"

type ProductionDetailRepository interface {
	Save(production model.ProductionDetail) error
	Update(production model.ProductionDetail) error
	FindById(productionId int) (*model.ProductionDetail, error)
	FindAll() (productions []model.ProductionDetail, err error)
	GetAllByQuery(production model.ProductionDetail) (productions []model.ProductionDetail, err error)
	GetOneByQuery(production model.ProductionDetail) (productionData model.ProductionDetail, err error)
}
