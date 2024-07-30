package repository

import "github.com/iniakunhuda/logistik-tani/inventory/model"

type InventoryRepository interface {
	Save(produk model.Produk) error
	Update(produk model.Produk) error
	Delete(produkId int) error
	FindById(produkId int) (*model.Produk, error)
	FindAll() (produks []model.Produk, err error)
	GetAllByQuery(produk model.Produk) (produks []model.Produk, err error)
	GetOneByQuery(produk model.Produk) (produkData model.Produk, err error)
}
