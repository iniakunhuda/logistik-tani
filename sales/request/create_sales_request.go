package request

import "time"

type CreateSalesRequest struct {
	IDSeller  int                      `json:"id_seller" validate:"required"` // ID pembibit
	IDBuyer   int                      `json:"id_buyer" validate:"required"`  // ID petani
	Products  []CreateSalesItemRequest `json:"products" validate:"required"`
	SalesDate time.Time                `json:"sales_date" validate:"required"`
}

type CreateSalesItemRequest struct {
	IDProduct int    `json:"id_product" validate:"required"`
	Category  string `json:"category" validate:"required"`
	Price     int    `json:"price" validate:"required"`
	Qty       int    `json:"qty" validate:"required"`
	Catatan   string `json:"catatan"`
}
