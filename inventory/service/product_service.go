package service

import (
	"github.com/iniakunhuda/logistik-tani/inventory/model"
	"github.com/iniakunhuda/logistik-tani/inventory/request"
	"github.com/iniakunhuda/logistik-tani/inventory/response"
)

type ProductService interface {
	Create(produk request.CreateProdukRequest) error
	FindById(produkOwnerId int, userId ...string) (response.ProductResponse, error)
	FindAll(produk *model.Product, userId ...string) ([]response.ProductResponse, error)
	Update(produkOwnerId int, produk request.UpdateProdukRequest) (response.ProductResponse, error)

	UpdateReduceStock(produkOwnerId int, stock int) error
	UpdateIncreaseStock(produkOwnerId int, stock int) error

	// ketika sales pembeli petani, maka produk petani akan otomatis dibuat
	AutoCreateProductPetani(produk request.CreateProdukRequest) error
}
