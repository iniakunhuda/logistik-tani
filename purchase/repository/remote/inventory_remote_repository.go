package remote

import (
	"github.com/iniakunhuda/logistik-tani/purchase/response"
	panenresponse "github.com/iniakunhuda/logistik-tani/purchase/response/panen_response"
)

type InventoryRemoteRepository interface {
	GetAll() ([]response.ProductResponse, error)
	GetDetail(id string) (response.ProductResponse, error)
	UpdateReduceStok(id string, stok string) error
	UpdateIncreaseStok(id string, stok string) error

	AutoCreateProdukPetani(produk response.ProductResponse, qty uint, idPembeli uint) error

	GetPanenAll() ([]panenresponse.ProductionRowResponse, error)
	GetPanenDetail(id string) (panenresponse.ProductionRowResponse, error)
}
