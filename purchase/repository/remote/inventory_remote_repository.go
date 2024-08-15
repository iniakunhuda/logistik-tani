package remote

import "github.com/iniakunhuda/logistik-tani/purchase/response"

type InventoryRemoteRepository interface {
	GetDetail(id string) (response.ProdukResponse, error)
	UpdateIncreaseStok(id string, stok string) error

	PostCreatePetani(produk response.ProdukResponse, qty uint, idPembeli uint) error
}
