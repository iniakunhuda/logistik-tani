package repository

import "github.com/iniakunhuda/logistik-tani/sales/model"

type SalesRepository interface {
	Save(sale model.Sales, salesDetail []model.SalesDetail) error
	Delete(saleId int) error
	FindById(saleId int) (*model.Sales, error)
	FindAll() (sales []model.Sales, err error)
	FindLastRow() (sale *model.Sales, err error)
}
