package request

import "time"

type CreateSalesIgmRequest struct {
	IDSeller     int                       `json:"id_seller" validate:"required"` // ID pembibit
	Product      CreateSalesIgmItemRequest `json:"product" validate:"required"`
	SalesDate    time.Time                 `json:"sales_date" validate:"required"`
	BuyerName    string                    `json:"buyer_name" validate:"required"`
	BuyerTelp    string                    `json:"buyer_telp" validate:"required"`
	BuyerAddress string                    `json:"buyer_address" validate:"required"`
}

type CreateSalesIgmItemRequest struct {
	IDProduct int    `json:"id_product" validate:"required"`
	Category  string `json:"category" validate:"required"`
	Price     int    `json:"price" validate:"required"`
	Qty       int    `json:"qty" validate:"required"`
	Catatan   string `json:"catatan"`
}
