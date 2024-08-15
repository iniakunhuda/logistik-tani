package repository

import "github.com/iniakunhuda/logistik-tani/sales/model"

type SalesDetailRepository interface {
	Save(sale model.SalesDetail) error
	Delete(saleId int) error
	FindById(saleId int) (*model.SalesDetail, error)
	FindAll() (sales []model.SalesDetail, err error)
	FindLastRow() (sale *model.SalesDetail, err error)
}
