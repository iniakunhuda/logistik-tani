package remote

import "github.com/iniakunhuda/logistik-tani/sales/response"

type InventoryRemoteRepository interface {
	GetAll() ([]response.ProductResponse, error)
	GetDetail(id string) (response.ProductResponse, error)
	UpdateReduceStok(id string, stok string) error
	UpdateIncreaseStok(id string, stok string) error

	AutoCreateProdukPetani(produk response.ProductResponse, qty uint, idPembeli uint) error
}
