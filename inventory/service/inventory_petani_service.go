package service

import (
	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/response"
)

type InventoryPetaniService interface {
	SetUserId(userId string)

	Create(produk request.CreateProdukRequest) error
	Update(produkId int, produk request.UpdateProdukRequest) error
	Delete(produkId int) error
	FindById(produkId int) (response.ProdukPetaniResponse, error)
	FindAll(produk *model.ProdukPetani) ([]response.ProdukPetaniResponse, error)

	UpdateReduceStock(produkId int, stock int) error
}
