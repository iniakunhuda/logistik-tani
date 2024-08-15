package remote

import "github.com/iniakunhuda/logistik-tani/sales/response"

type InventoryRemoteRepository interface {
	GetDetail(id string) (response.ProdukResponse, error)
	UpdateReduceStok(id string, stok string) error

	PostCreatePetani(produk response.ProdukResponse, qty uint, idPembeli uint) error
}
