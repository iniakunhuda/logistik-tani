package repository

import (
	"github.com/iniakunhuda/logistik-tani/purchase/model"
)

type PurchaseDetailRepository interface {
	Save(sale model.PurchaseDetail) error
	Delete(saleId int) error
	FindById(saleId int) (*model.PurchaseDetail, error)
	FindAll() (sales []model.PurchaseDetail, err error)
	FindLastRow() (sale *model.PurchaseDetail, err error)
}
