package repository

import "github.com/iniakunhuda/logistik-tani/inventory/model"

type InventoryPetaniRepository interface {
	Save(produk model.ProdukPetani) error
	Update(produk model.ProdukPetani) error
	Delete(produkId int) error
	FindById(produkId int) (*model.ProdukPetani, error)
	FindAll() (produks []model.ProdukPetani, err error)
	GetAllByQuery(produk model.ProdukPetani) (produks []model.ProdukPetani, err error)
	GetOneByQuery(produk model.ProdukPetani) (produkData model.ProdukPetani, err error)
}
