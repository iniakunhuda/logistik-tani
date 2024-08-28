package repository

import "github.com/iniakunhuda/logistik-tani/inventory/model"

type ProductRepository interface {
	Save(product model.Product) error
	Update(product model.Product) error
	Delete(productId int) error
	FindById(productId int) (*model.Product, error)
	FindAll() (products []model.Product, err error)
	FindByName(productName string) (*model.Product, error)
	GetAllByQuery(product model.Product) (products []model.Product, err error)
	GetOneByQuery(product model.Product) (productData model.Product, err error)
}
