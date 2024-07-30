package repository

import "github.com/iniakunhuda/logistik-tani/sales/model"

type SalesRepository interface {
	Save(sale model.Sales) error
	Update(sale model.Sales) error
	Delete(saleId int) error
	FindById(saleId int) (*model.Sales, error)
	FindAll() (sales []model.Sales, err error)
	GetAllByQuery(sale model.Sales) (sales []model.Sales, err error)
	GetOneByQuery(sale model.Sales) (saleData model.Sales, err error)
}
