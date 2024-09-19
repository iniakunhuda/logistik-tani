package request

import "time"

type CreatePurchaseRequest struct {
	IDSeller      int                         `json:"id_seller"`
	SellerName    string                      `json:"seller_name"`
	SellerAddress string                      `json:"seller_address"`
	SellerTelp    string                      `json:"seller_telp"`
	IDBuyer       uint                        `json:"id_buyer" validate:"required"`
	Products      []CreatePurchaseItemRequest `json:"products" validate:"required"`
	PurchaseDate  time.Time                   `json:"purchase_date" validate:"required"`
}

type CreatePurchaseItemRequest struct {
	IDProduct int    `json:"id_product" validate:"required"`
	Category  string `json:"category" validate:"required"`
	Price     int    `json:"price" validate:"required"`
	Qty       int    `json:"qty" validate:"required"`
	Catatan   string `json:"catatan"`
}
