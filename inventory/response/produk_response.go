package response

import "github.com/iniakunhuda/logistik-tani/inventory/model"

type ProdukResponse struct {
	model.Produk
	Password *struct{} `json:"password,omitempty"`
}
