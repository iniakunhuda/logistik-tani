package response

import "github.com/iniakunhuda/logistik-tani/purchase/model"

type PurchaseResponse struct {
	model.Purchase
	IDSeller      uint         `json:"id_seller"`
	SellerName    string       `json:"seller_name"`
	SellerAddress string       `json:"seller_address"`
	SellerTelp    string       `json:"seller_telp"`
	BuyerDetail   UserResponse `json:"buyer_detail"`
}
