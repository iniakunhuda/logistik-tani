package response

import "github.com/iniakunhuda/logistik-tani/sales/model"

type SalesResponse struct {
	model.Sales
	PenjualDetail UserResponse `json:"penjual_detail"`
	PembeliDetail UserResponse `json:"pembeli_detail"`
}
