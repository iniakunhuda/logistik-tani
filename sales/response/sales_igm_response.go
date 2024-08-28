package response

import "github.com/iniakunhuda/logistik-tani/sales/model"

type SalesIgmResponse struct {
	model.SalesIgm
	SellerDetail UserResponse `json:"seller_detail"`
}
