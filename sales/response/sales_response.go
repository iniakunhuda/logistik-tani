package response

import "github.com/iniakunhuda/logistik-tani/sales/model"

type SalesResponse struct {
	model.Sales
	SellerDetail UserResponse `json:"seller_detail"`
	BuyerDetail  UserResponse `json:"buyer_detail"`
}
