package repository

import "github.com/iniakunhuda/logistik-tani/sales/model"

type SalesIgmRepository interface {
	Save(sale model.SalesIgm) error
	Delete(saleId int) error
	Update(sale model.SalesIgm) error
	FindById(saleId int) (*model.SalesIgm, error)
	FindAll() (sales []model.SalesIgm, err error)
	FindLastRow() (sale *model.SalesIgm, err error)
	GetAllByQuery(sale model.SalesIgm) (sales []model.SalesIgm, err error)
	GetOneByQuery(sale model.SalesIgm) (saleData model.SalesIgm, err error)
}
