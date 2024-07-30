package service

import (
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/response"
)

type InventoryService interface {
	Create(produk request.CreateProdukRequest) error
	Update(produkId int, produk request.UpdateUserRequest) error
	Delete(produkId int) error
	FindById(produkId int) (response.ProdukResponse, error)
	FindAll() ([]response.ProdukResponse, error)
}
