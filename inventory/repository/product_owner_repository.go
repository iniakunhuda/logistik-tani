package repository

import "github.com/iniakunhuda/logistik-tani/inventory/model"

type ProductOwnerRepository interface {
	Save(product model.ProductOwner) error
	Update(product model.ProductOwner) error
	Delete(productId int) error
	FindById(productId int) (*model.ProductOwner, error)
	FindAll() (products []model.ProductOwner, err error)
	GetAllByQuery(product model.ProductOwner) (products []model.ProductOwner, err error)
	GetOneByQuery(product model.ProductOwner) (productData model.ProductOwner, err error)

	GetAllByProduk(product model.Product, idUser ...string) (products []model.ProductOwner, err error)
}
